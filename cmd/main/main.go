package main

import (
	"encoding/json"
	"go-api-with-ratelimit/internal/middleware"
	"go-api-with-ratelimit/internal/mux"
	"go-api-with-ratelimit/internal/ratelimiter"
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	mux := mux.New()

	/*
		tambah 2 token/detik ke bucket
		burst yg diperbolehkan 10 token
	*/
	rl := ratelimiter.New(rate.Limit(2), 10)

	m := middleware.New(rl)

	mux.RegisterMiddleware(m.RateLimiterMiddleware)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := Response{Message: "Hello, World"}
		json.NewEncoder(w).Encode(response)
	})

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = mux

	log.Println("Server started at localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
