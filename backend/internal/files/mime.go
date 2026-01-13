package files

import (
	"path/filepath"
	"strings"
)

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
		".zip": "application/zip",
		".tar": "application/x-tar",
		".gz":  "application/gzip",
		".rar": "application/x-rar-compressed",
		".7z":  "application/x-7z-compressed",

		// Audio/Video
		".mp3": "audio/mpeg",
		".wav": "audio/wav",
		".mp4": "video/mp4",
		".avi": "video/x-msvideo",
		".mov": "video/quicktime",

		// Executables
		".exe": "application/x-msdownload",
		".msi": "application/x-msi",
		".deb": "application/x-debian-package",
		".rpm": "application/x-rpm",
		".dmg": "application/x-apple-diskimage",
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

// isInlineMimeType checks if a MIME type should be displayed inline in browser
func isInlineMimeType(mimeType string) bool {
	inlineTypes := []string{
		"image/",
		"video/",
		"audio/",
		"application/pdf",
		"text/",
		"application/json",
		"application/xml",
		"application/javascript",
	}

	for _, inlineType := range inlineTypes {
		if strings.HasPrefix(mimeType, inlineType) {
			return true
		}
	}

	return false
}