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
	mux.HandleFunc("GET /details/", handler.getFileDetails)
	mux.HandleFunc("DELETE /", handler.deleteFile)

	return mux
}

func (h *FileHandler) getFileList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("file handler get file list")
	dirList, err := h.svc.List()
	if err != nil {
		fmt.Fprint(w, "file read error")
	}

	fmt.Fprint(w, dirList)
}

func (h *FileHandler) getFileDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("file handler get file details")
	h.svc.ReadFile()
}

func (h *FileHandler) deleteFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("file handler delete file")

	query := r.URL.Query()
	targetFile := query.Get("targetFile")

	if targetFile == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "target file not provineded")
		return
	}

	h.svc.DeleteFile(targetFile)
}
