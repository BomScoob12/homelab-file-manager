package routes

import (
	"net/http"

	"github.com/BomScoob12/homelab-file-manager/internal/files"
)

func NewRouter() http.Handler {
	// mux = multiplexter (router)
	mux := http.NewServeMux()

	// API routes
	mux.Handle("/file/", http.StripPrefix("/file", files.NewHandler()))

	// Serve frontend static files (for production)
	// Uncomment this when you build the frontend
	// fs := http.FileServer(http.Dir("../frontend/dist/"))
	// mux.Handle("/", fs)

	// For development, you can serve a simple index page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
    <title>File Manager API</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 50px auto; padding: 20px; }
        .endpoint { background: #f5f5f5; padding: 10px; margin: 10px 0; border-radius: 5px; }
        .method { color: #007acc; font-weight: bold; }
    </style>
</head>
<body>
    <h1>File Manager API</h1>
    <p>Backend is running! The Vue.js frontend should be running on <a href="http://localhost:3000">http://localhost:3000</a></p>
    
    <h2>Available API Endpoints:</h2>
    <div class="endpoint">
        <span class="method">GET</span> /file/list?path=/ - List files and directories
    </div>
    <div class="endpoint">
        <span class="method">GET</span> /file/details?path=/file.txt - Get file details
    </div>
    <div class="endpoint">
        <span class="method">GET</span> /file/open?path=/file.txt - Read file content
    </div>
    <div class="endpoint">
        <span class="method">DELETE</span> /file/delete?path=/file.txt - Delete file
    </div>
    
    <h2>Quick Test:</h2>
    <p><a href="/file/list?path=/">List root directory</a></p>
</body>
</html>
			`))
		} else {
			http.NotFound(w, r)
		}
	})

	return mux
}
