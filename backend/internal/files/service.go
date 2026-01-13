package files

import (
	"fmt"
	"path/filepath"
	"strings"
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
	basePath := "/WorkDir" // Default for Docker/Unix
	
	// For Windows development, use C:\tmp\WorkDir
	if filepath.Separator == '\\' {
		basePath = `C:\tmp\WorkDir`
	}
	
	return &FileService{
		basePath: basePath,
		fsUtils:  fs.NewFileSystemUtils(),
	}
}

// Response models
type FileItem struct {
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	IsDir       bool      `json:"isDir"`
	FileType    string    `json:"fileType"`
	Size        int64     `json:"size"`
	ModTime     time.Time `json:"modTime"`
	Permissions string    `json:"permissions"`
	Extension   string    `json:"extension,omitempty"`
}

type FileListResponse struct {
	Success     bool       `json:"success"`
	Path        string     `json:"path"`
	Items       []FileItem `json:"items"`
	TotalItems  int        `json:"totalItems"`
	TotalSize   int64      `json:"totalSize"`
	RequestTime time.Time  `json:"requestTime"`
}

type FileDetailsResponse struct {
	Success     bool      `json:"success"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	FullPath    string    `json:"fullPath"`
	IsDir       bool      `json:"isDir"`
	Size        int64     `json:"size"`
	ModTime     time.Time `json:"modTime"`
	MimeType    string    `json:"mimeType"`
	Permissions string    `json:"permissions"`
	Extension   string    `json:"extension,omitempty"`
	RequestTime time.Time `json:"requestTime"`
}

type FileContentResponse struct {
	Success     bool      `json:"success"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Content     string    `json:"content"`
	Size        int64     `json:"size"`
	MimeType    string    `json:"mimeType"`
	Encoding    string    `json:"encoding"`
	RequestTime time.Time `json:"requestTime"`
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
		fileItem := FileItem{
			Name:        entry.Name(),
			Path:        itemPath,
			IsDir:       entry.IsDir(),
			FileType:    entry.Type().String(),
			Size:        info.Size(),
			ModTime:     info.ModTime(),
			Permissions: info.Mode().String(),
			Extension:   filepath.Ext(entry.Name()),
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

// Helper methods for the service

// validateAndConstructPath validates the path and constructs the full system path
func (s *FileService) validateAndConstructPath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("path cannot be empty")
	}

	// Clean the path
	cleanPath := filepath.Clean(path)
	
	// Construct full path
	fullPath := filepath.Join(s.basePath, cleanPath)
	
	// Ensure the path doesn't escape the base directory
	if !strings.HasPrefix(filepath.Clean(fullPath), s.basePath) {
		return "", fmt.Errorf("invalid path: access denied - path escapes base directory")
	}

	return fullPath, nil
}

// getMimeType returns a MIME type based on file extension
func getMimeType(filePath string) string {
	ext := strings.ToLower(filepath.Ext(filePath))
	
	mimeTypes := map[string]string{
		// Text files
		".txt":  "text/plain",
		".md":   "text/markdown",
		".csv":  "text/csv",
		".log":  "text/plain",
		".conf": "text/plain",
		".cfg":  "text/plain",
		".ini":  "text/plain",
		
		// Code files
		".go":   "text/x-go",
		".js":   "application/javascript",
		".ts":   "application/typescript",
		".py":   "text/x-python",
		".java": "text/x-java-source",
		".c":    "text/x-c",
		".cpp":  "text/x-c++",
		".h":    "text/x-c",
		".php":  "application/x-httpd-php",
		".rb":   "text/x-ruby",
		".sh":   "application/x-sh",
		".bat":  "application/x-bat",
		".ps1":  "application/x-powershell",
		
		// Web files
		".html": "text/html",
		".htm":  "text/html",
		".css":  "text/css",
		".scss": "text/x-scss",
		".sass": "text/x-sass",
		".less": "text/x-less",
		
		// Data files
		".json": "application/json",
		".xml":  "application/xml",
		".yaml": "application/x-yaml",
		".yml":  "application/x-yaml",
		".toml": "application/toml",
		
		// Images
		".png":  "image/png",
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".gif":  "image/gif",
		".bmp":  "image/bmp",
		".svg":  "image/svg+xml",
		".webp": "image/webp",
		".ico":  "image/x-icon",
		
		// Documents
		".pdf":  "application/pdf",
		".doc":  "application/msword",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		".xls":  "application/vnd.ms-excel",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		".ppt":  "application/vnd.ms-powerpoint",
		".pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		
		// Archives
		".zip":  "application/zip",
		".tar":  "application/x-tar",
		".gz":   "application/gzip",
		".rar":  "application/x-rar-compressed",
		".7z":   "application/x-7z-compressed",
		
		// Audio/Video
		".mp3":  "audio/mpeg",
		".wav":  "audio/wav",
		".mp4":  "video/mp4",
		".avi":  "video/x-msvideo",
		".mov":  "video/quicktime",
		
		// Executables
		".exe":  "application/x-msdownload",
		".msi":  "application/x-msi",
		".deb":  "application/x-debian-package",
		".rpm":  "application/x-rpm",
		".dmg":  "application/x-apple-diskimage",
	}

	if mimeType, exists := mimeTypes[ext]; exists {
		return mimeType
	}
	
	return "application/octet-stream"
}

// isBinaryMimeType checks if a MIME type represents binary content
func isBinaryMimeType(mimeType string) bool {
	textTypes := []string{
		"text/",
		"application/json",
		"application/xml",
		"application/javascript",
		"application/typescript",
		"application/x-yaml",
		"application/toml",
	}
	
	for _, textType := range textTypes {
		if strings.HasPrefix(mimeType, textType) {
			return false
		}
	}
	
	return true
}