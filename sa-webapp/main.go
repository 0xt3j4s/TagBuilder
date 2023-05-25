package main

import (
	"fmt"
	"log"
	"net/http"
	"tagbuilder/handlers"
)

func main() {
	fmt.Println("Hello, TagBuilder!")

	// Call the buildTagsHandler function
	tagsHandler := handlers.BuildTagsHandler()

	// Create a server and register the tagsHandler
	server := &http.Server{
		Addr:    ":8081",
		Handler: corsMiddleware(tagsHandler),
	}

	// Start the server
	log.Fatal(server.ListenAndServe())
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}