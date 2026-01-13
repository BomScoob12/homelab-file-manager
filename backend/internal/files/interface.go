package files

import "net/http"

// FileServiceInterface defines the contract for file service operations
type FileServiceInterface interface {
	ListFiles(path string) (*FileListResponse, error)
	GetFileDetails(path string) (*FileDetailsResponse, error)
	DeleteFile(path string) error
	OpenFile(path string) (*FileContentResponse, error)
	ServeRawFile(w http.ResponseWriter, path string) error
}