package files

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

// FileHandler handles HTTP requests for file operations
type FileHandler struct {
	svc FileServiceInterface
}

// FileServiceInterface defines the contract for file service operations
type FileServiceInterface interface {
	ListFiles(path string) (*FileListResponse, error)
	GetFileDetails(path string) (*FileDetailsResponse, error)
	DeleteFile(path string) error
	OpenFile(path string) (*FileContentResponse, error)
}

// NewHandler creates a new file handler with proper routing
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	handler := &FileHandler{
		svc: NewFileService(),
	}

	// File management endpoints
	mux.HandleFunc("/list", handler.handleListFiles)
	mux.HandleFunc("/open", handler.handleOpenFile)
	mux.HandleFunc("/details", handler.handleGetFileDetails)
	mux.HandleFunc("/delete", handler.handleDeleteFile)

	// Add middleware for logging and CORS
	return handler.withMiddleware(mux)
}

// handleListFiles handles GET /file/list - Lists files and directories
func (h *FileHandler) handleListFiles(w http.ResponseWriter, r *http.Request) {
	// Check HTTP method
	if r.Method != http.MethodGet {
		h.sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract and validate path parameter
	path := r.URL.Query().Get("path")
	if path == "" {
		path = "/"
	}

	// Clean and validate path
	cleanPath := filepath.Clean(path)

	// Debug logging
	log.Printf("Received path: %q, cleaned path: %q", path, cleanPath)

	if !isValidPath(cleanPath) {
		h.sendErrorResponse(w, "Invalid path provided", http.StatusBadRequest)
		return
	}

	// Call service layer
	result, err := h.svc.ListFiles(cleanPath)
	if err != nil {
		log.Printf("Error listing files for path %s: %v", cleanPath, err)
		h.handleServiceError(w, err)
		return
	}

	// Send successful response
	h.sendJSONResponse(w, result, http.StatusOK)
}

// handleOpenFile handles GET /file/open - Opens and reads file content
func (h *FileHandler) handleOpenFile(w http.ResponseWriter, r *http.Request) {
	// Check HTTP method
	if r.Method != http.MethodGet {
		h.sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract and validate file path
	filePath := r.URL.Query().Get("path")
	if filePath == "" {
		h.sendErrorResponse(w, "File path is required", http.StatusBadRequest)
		return
	}

	// Clean and validate path
	cleanPath := filepath.Clean(filePath)
	if !isValidPath(cleanPath) {
		h.sendErrorResponse(w, "Invalid file path provided", http.StatusBadRequest)
		return
	}

	// Call service layer
	result, err := h.svc.OpenFile(cleanPath)
	if err != nil {
		log.Printf("Error opening file %s: %v", cleanPath, err)
		h.handleServiceError(w, err)
		return
	}

	// Send successful response
	h.sendJSONResponse(w, result, http.StatusOK)
}

// handleGetFileDetails handles GET /file/details - Gets detailed file information
func (h *FileHandler) handleGetFileDetails(w http.ResponseWriter, r *http.Request) {
	// Check HTTP method
	if r.Method != http.MethodGet {
		h.sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract and validate file path
	filePath := r.URL.Query().Get("path")
	if filePath == "" {
		h.sendErrorResponse(w, "File path is required", http.StatusBadRequest)
		return
	}

	// Clean and validate path
	cleanPath := filepath.Clean(filePath)
	if !isValidPath(cleanPath) {
		h.sendErrorResponse(w, "Invalid file path provided", http.StatusBadRequest)
		return
	}

	// Call service layer
	result, err := h.svc.GetFileDetails(cleanPath)
	if err != nil {
		log.Printf("Error getting file details for %s: %v", cleanPath, err)
		h.handleServiceError(w, err)
		return
	}

	// Send successful response
	h.sendJSONResponse(w, result, http.StatusOK)
}

// handleDeleteFile handles DELETE /file/delete - Deletes a file
func (h *FileHandler) handleDeleteFile(w http.ResponseWriter, r *http.Request) {
	// Check HTTP method
	if r.Method != http.MethodDelete {
		h.sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract and validate file path
	filePath := r.URL.Query().Get("path")
	if filePath == "" {
		h.sendErrorResponse(w, "File path is required", http.StatusBadRequest)
		return
	}

	// Clean and validate path
	cleanPath := filepath.Clean(filePath)
	if !isValidPath(cleanPath) {
		h.sendErrorResponse(w, "Invalid file path provided", http.StatusBadRequest)
		return
	}

	// Call service layer
	err := h.svc.DeleteFile(cleanPath)
	if err != nil {
		log.Printf("Error deleting file %s: %v", cleanPath, err)
		h.handleServiceError(w, err)
		return
	}

	// Send successful response
	response := map[string]interface{}{
		"success": true,
		"message": "File deleted successfully",
		"path":    cleanPath,
	}
	h.sendJSONResponse(w, response, http.StatusOK)
}

// Helper methods for the handler

// sendJSONResponse sends a JSON response with proper headers
func (h *FileHandler) sendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// sendErrorResponse sends an error response with proper format
func (h *FileHandler) sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	errorResponse := map[string]interface{}{
		"success": false,
		"error":   message,
		"code":    statusCode,
	}
	h.sendJSONResponse(w, errorResponse, statusCode)
}

// handleServiceError handles errors from the service layer
func (h *FileHandler) handleServiceError(w http.ResponseWriter, err error) {
	if strings.Contains(err.Error(), "no such file") || strings.Contains(err.Error(), "not found") {
		h.sendErrorResponse(w, "File or directory not found", http.StatusNotFound)
	} else if strings.Contains(err.Error(), "access denied") || strings.Contains(err.Error(), "permission denied") {
		h.sendErrorResponse(w, "Access denied", http.StatusForbidden)
	} else if strings.Contains(err.Error(), "invalid path") {
		h.sendErrorResponse(w, "Invalid path provided", http.StatusBadRequest)
	} else {
		h.sendErrorResponse(w, "Internal server error", http.StatusInternalServerError)
	}
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

// withMiddleware adds middleware to the handler
func (h *FileHandler) withMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Log request
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Call next handler
		next.ServeHTTP(w, r)

		// Log completion
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}
