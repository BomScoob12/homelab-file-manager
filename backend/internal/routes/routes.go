package routes

import (
	"net/http"

	"github.com/BomScoob12/homelab-file-manager/internal/files"
)

func NewRouter() http.Handler {
	// mux = multiplexter (router)
	mux := http.NewServeMux()

	mux.Handle("/files/", http.StripPrefix("/files", files.NewHandler()))

	return mux
}
