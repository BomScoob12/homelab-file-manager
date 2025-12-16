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
)

func main() {

	port := ":8080"
	// mux = multiplexter (router)
	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	go func() {
		fmt.Println("server running at port", port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("listen error %v", err)
		}
	}()

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

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
