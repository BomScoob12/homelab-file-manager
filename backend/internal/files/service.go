package files

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/BomScoob12/homelab-file-manager/internal/fs"
)

// FileService implements file management operations
type FileService struct {
	basePath string
	fsUtils  fs.FileSystemInterface
}

// NewFileService creates a new file service instance
func NewFileService() *FileService {
	// Get base path from environment variable, with fallback
	basePath := os.Getenv("FILE_MANAGER_BASE_PATH")
	if basePath == "" {
		// Default to Docker mount point
		basePath = "/data"
	}

	return &FileService{
		basePath: basePath,
		fsUtils:  fs.NewFileSystemUtils(),
	}
}

// ListFiles lists all files and directories in the specified path
func (s *FileService) ListFiles(path string) (*FileListResponse, error) {
	// Validate and construct full path
	fullPath, err := s.validateAndConstructPath(path)
	if err != nil {
		return nil, fmt.Errorf("path validation failed: %w", err)
	}

	// Check if directory exists and is accessible
	if !s.fsUtils.IsDirectory(fullPath) {
		return nil, fmt.Errorf("path is not a directory or does not exist: %s", path)
	}

	// List directory contents
	entries, err := s.fsUtils.ListDirectory(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to list directory: %w", err)
	}

	// Process entries
	var fileItems []FileItem
	var totalSize int64

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue // Skip files we can't get info for
		}

		itemPath := filepath.Join(path, entry.Name())
		fullItemPath := filepath.Join(fullPath, entry.Name())
		
		var mimeType string
		if entry.IsDir() {
			mimeType = "inode/directory"
		} else {
			mimeType = getMimeType(fullItemPath)
		}
		
		fileItem := FileItem{
			Name:        entry.Name(),
			Path:        itemPath,
			IsDir:       entry.IsDir(),
			FileType:    entry.Type().String(),
			Size:        info.Size(),
			ModTime:     info.ModTime(),
			Permissions: info.Mode().String(),
			Extension:   filepath.Ext(entry.Name()),
			MimeType:    mimeType,
		}

		fileItems = append(fileItems, fileItem)
		totalSize += info.Size()
	}

	return &FileListResponse{
		Success:     true,
		Path:        path,
		Items:       fileItems,
		TotalItems:  len(fileItems),
		TotalSize:   totalSize,
		RequestTime: time.Now(),
	}, nil
}

// GetFileDetails gets detailed information about a specific file or directory
func (s *FileService) GetFileDetails(filePath string) (*FileDetailsResponse, error) {
	// Validate and construct full path
	fullPath, err := s.validateAndConstructPath(filePath)
	if err != nil {
		return nil, fmt.Errorf("path validation failed: %w", err)
	}

	// Get file information
	info, err := s.fsUtils.GetFileInfo(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	return &FileDetailsResponse{
		Success:     true,
		Name:        info.Name(),
		Path:        filePath,
		FullPath:    fullPath,
		IsDir:       info.IsDir(),
		Size:        info.Size(),
		ModTime:     info.ModTime(),
		MimeType:    getMimeType(fullPath),
		Permissions: info.Mode().String(),
		Extension:   filepath.Ext(info.Name()),
		RequestTime: time.Now(),
	}, nil
}

// OpenFile opens and reads the content of a file
func (s *FileService) OpenFile(filePath string) (*FileContentResponse, error) {
	// Validate and construct full path
	fullPath, err := s.validateAndConstructPath(filePath)
	if err != nil {
		return nil, fmt.Errorf("path validation failed: %w", err)
	}

	// Check if it's a file (not directory)
	info, err := s.fsUtils.GetFileInfo(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	if info.IsDir() {
		return nil, fmt.Errorf("cannot open directory as file: %s", filePath)
	}

	// Check file size limit (10MB)
	const maxFileSize = 10 * 1024 * 1024
	if info.Size() > maxFileSize {
		return nil, fmt.Errorf("file too large to open: %d bytes (max: %d bytes)", info.Size(), maxFileSize)
	}

	// Read file content
	content, err := s.fsUtils.ReadFileContent(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %w", err)
	}

	mimeType := getMimeType(fullPath)
	encoding := "utf-8"

	// Check if it's a binary file
	if isBinaryMimeType(mimeType) {
		encoding = "binary"
		// For binary files, we might want to return base64 encoded content
		// content = base64.StdEncoding.EncodeToString([]byte(content))
	}

	return &FileContentResponse{
		Success:     true,
		Name:        info.Name(),
		Path:        filePath,
		Content:     content,
		Size:        info.Size(),
		MimeType:    mimeType,
		Encoding:    encoding,
		RequestTime: time.Now(),
	}, nil
}

// DeleteFile deletes a file or directory
func (s *FileService) DeleteFile(targetPath string) error {
	// Validate and construct full path
	fullPath, err := s.validateAndConstructPath(targetPath)
	if err != nil {
		return fmt.Errorf("path validation failed: %w", err)
	}

	// Check if file exists
	if !s.fsUtils.Exists(fullPath) {
		return fmt.Errorf("file or directory not found: %s", targetPath)
	}

	// Delete the file or directory
	err = s.fsUtils.Delete(fullPath)
	if err != nil {
		return fmt.Errorf("failed to delete: %w", err)
	}

	return nil
}

// ServeRawFile serves raw file content directly (for images, PDFs, etc.)
func (s *FileService) ServeRawFile(w http.ResponseWriter, filePath string) error {
	// Validate and construct full path
	fullPath, err := s.validateAndConstructPath(filePath)
	if err != nil {
		return fmt.Errorf("path validation failed: %w", err)
	}

	// Check if file exists and is not a directory
	info, err := s.fsUtils.GetFileInfo(fullPath)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	if info.IsDir() {
		return fmt.Errorf("cannot serve directory as file: %s", filePath)
	}

	// Open the file
	file, err := os.Open(fullPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Set appropriate headers
	mimeType := getMimeType(fullPath)
	w.Header().Set("Content-Type", mimeType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size()))
	
	// Set cache headers for better performance
	w.Header().Set("Cache-Control", "public, max-age=3600")
	w.Header().Set("Last-Modified", info.ModTime().UTC().Format(http.TimeFormat))

	// For downloads, set Content-Disposition header
	if !isInlineMimeType(mimeType) {
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", info.Name()))
	}

	// Copy file content to response
	_, err = io.Copy(w, file)
	if err != nil {
		return fmt.Errorf("failed to serve file content: %w", err)
	}

	return nil
}
