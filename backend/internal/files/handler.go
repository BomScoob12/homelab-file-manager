package files

import (
	"fmt"
	"net/http"
)

type FileHandler struct {
	svc *FileService
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	handler := FileHandler{
		svc: &FileService{},
	}

	mux.HandleFunc("GET /", handler.getFileList)
	mux.HandleFunc("DELETE /", handler.deleteFile)

	return mux
}

func (h *FileHandler) getFileList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("file handler get file list")
	h.svc.List()
}

func (h *FileHandler) deleteFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("file handler delete file")
	h.svc.DeleteFile()
}
