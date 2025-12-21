package routes

import (
	"net/http"

	"github.com/BomScoob12/homelab-file-manager/internal/files"
)

func NewRouter() http.Handler {
	// mux = multiplexter (router)
	mux := http.NewServeMux()

	mux.Handle("/file/", http.StripPrefix("/file", files.NewHandler()))

	return mux
}
