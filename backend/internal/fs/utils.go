package fs

import (
	"fmt"
	"io"
	"os"
)

// FileSystemInterface defines the contract for file system operations
type FileSystemInterface interface {
	ReadFileContent(path string) (string, error)
	ListDirectory(path string) ([]os.DirEntry, error)
	GetFileInfo(path string) (os.FileInfo, error)
	IsDirectory(path string) bool
	Exists(path string) bool
	Delete(path string) error
}

// FileSystemUtils implements FileSystemInterface
type FileSystemUtils struct{}

// NewFileSystemUtils creates a new FileSystemUtils instance
func NewFileSystemUtils() *FileSystemUtils {
	return &FileSystemUtils{}
}

// ReadFileContent reads the entire content of a file
func (fs *FileSystemUtils) ReadFileContent(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}

// ListDirectory lists all entries in a directory
func (fs *FileSystemUtils) ListDirectory(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}
	return entries, nil
}

// GetFileInfo gets file information
func (fs *FileSystemUtils) GetFileInfo(path string) (os.FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}
	return info, nil
}

// IsDirectory checks if the path is a directory
func (fs *FileSystemUtils) IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// Exists checks if a file or directory exists
func (fs *FileSystemUtils) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Delete removes a file or directory
func (fs *FileSystemUtils) Delete(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	if info.IsDir() {
		// Remove directory and all its contents
		err = os.RemoveAll(path)
	} else {
		// Remove single file
		err = os.Remove(path)
	}

	if err != nil {
		return fmt.Errorf("failed to delete: %w", err)
	}

	return nil
}
