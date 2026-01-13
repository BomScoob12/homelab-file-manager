package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BomScoob12/homelab-file-manager/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or could not be loaded: %v", err)
	}

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Get host from environment variable
	host := os.Getenv("HOST")

	// Get base path for logging
	basePath := os.Getenv("FILE_MANAGER_BASE_PATH")
	if basePath == "" {
		basePath = "default system path"
	}

	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: routes.NewRouter(),
	}

	go func() {
		fmt.Printf("🚀 File Manager Server starting...\n")
		fmt.Printf("📡 Server running at http://localhost:%s\n", port)
		fmt.Printf("📁 Base path: %s\n", basePath)
		fmt.Printf("🔧 Press Ctrl+C to stop\n\n")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server listen error: %v", err)
		}
	}()

	handleStopProcess(server)
}

func handleStopProcess(server *http.Server) {
	// stop signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("🛑 Shutdown signal received")

	// get process time from bg process + timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("❌ Graceful shutdown failed: %v", err)
	} else {
		log.Println("✅ Server stopped successfully")
	}
}
