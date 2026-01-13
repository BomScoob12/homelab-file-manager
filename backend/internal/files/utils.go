package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// validateAndConstructPath validates the path and constructs the full system path
func (s *FileService) validateAndConstructPath(path string) (string, error) {
	// Handle empty path as root
	if path == "" {
		return s.basePath, nil
	}

	// Handle root path
	if path == "/" {
		return s.basePath, nil
	}

	// Clean the path
	cleanPath := filepath.Clean(path)

	// If cleanPath is ".", it means we're at root
	if cleanPath == "." {
		return s.basePath, nil
	}

	// Construct full path
	fullPath := filepath.Join(s.basePath, cleanPath)

	// Clean both paths for comparison
	cleanFullPath := filepath.Clean(fullPath)
	cleanBasePath := filepath.Clean(s.basePath)

	// Ensure the path doesn't escape the base directory
	if !strings.HasPrefix(cleanFullPath, cleanBasePath) {
		return "", fmt.Errorf("invalid path: access denied - path escapes base directory")
	}

	return cleanFullPath, nil
}

// isValidPath validates if the path is safe to use
func isValidPath(path string) bool {
	// Check for empty path
	if path == "" {
		return true // Empty path is valid (represents root)
	}

	// Handle root paths
	if path == "/" || path == "." {
		return true
	}

	// Check for dangerous patterns
	dangerousPatterns := []string{"../", "..", "~"}
	for _, pattern := range dangerousPatterns {
		if strings.Contains(path, pattern) {
			return false
		}
	}

	return true
}