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
		Addr:    ":8080",
		Handler: tagsHandler,
	}

	// Start the server
	log.Fatal(server.ListenAndServe())
}
