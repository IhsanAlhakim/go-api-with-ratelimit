package main

import (
	"go-api-with-ratelimit/internal/mux"
	"log"
	"net/http"
)

func main() {
	mux := mux.New()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = mux

	log.Println("Server started at localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
