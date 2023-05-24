package handlers

import (
	"encoding/json"
	"bytes"
	"log"
	"net/http"
	"tagbuilder/dto"
)

func BuildTagsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body into a SentenceDto
		var sentenceDto dto.SentenceDto
		err := json.NewDecoder(r.Body).Decode(&sentenceDto)
		if err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Send a request to the Python application
		pythonURL := "http://localhost:5000/generate-tags/"
		reqBody, err := json.Marshal(sentenceDto)
		if err != nil {
			log.Println("Error marshaling request body:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		resp, err := http.Post(pythonURL, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Println("Error sending request to Python application:", err)
			http.Error(w, "Failed to fetch tags", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Parse the response body into a TagResponse
		var tagResponse dto.TagDto
		err = json.NewDecoder(resp.Body).Decode(&tagResponse)
		if err != nil {
			log.Println("Error decoding response body:", err)
			http.Error(w, "Failed to parse tags response", http.StatusInternalServerError)
			return
		}

		// Write the tags as a JSON response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(tagResponse)
		if err != nil {
			log.Println("Error encoding response:", err)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	})
}
