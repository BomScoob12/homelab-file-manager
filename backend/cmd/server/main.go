package main

import (
	"fmt"
	"net/http"
)

func main() {

	port := ":8080"
	// mux = multiplexter (router)
	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)

	server := &http.Server{
		Addr: port,
		Handler: mux,
	}
	
	fmt.Println("server running at port", port)
	server.ListenAndServe()
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
