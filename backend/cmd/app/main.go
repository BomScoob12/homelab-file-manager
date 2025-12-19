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
)

func main() {

	port := ":8080"
	server := &http.Server{
		Addr:    port,
		Handler: routes.NewRouter(),
	}

	go func() {
		fmt.Println("server running at port", port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("listen error %v", err)
		}
	}()

	handleStopProcess(server)
}

func handleStopProcess(server *http.Server) {
	// stop signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("shutdown signal received")

	// get process time from bg process + timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	} else {
		log.Println("server stopped successfully")
	}
}
