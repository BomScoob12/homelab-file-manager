package files

import "time"

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
	MimeType    string    `json:"mimeType,omitempty"`
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