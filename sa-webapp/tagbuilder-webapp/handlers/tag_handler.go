package handlers

import (
	"encoding/json"

	"net/http"
	"tagbuilder/dto"
)

func BuildTagsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body
		var requestData struct {
			Sentence string `json:"sentence"`
		}
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Generate tags based on the sentence
		tags := generateTags(requestData.Sentence)

		// Create a response with the generated tags
		response := dto.TagDto{
			Tags: tags,
		}

		// Convert the response to JSON
		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Set the response headers and write the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})
}

func generateTags(sentence string) []string {
	// Implement your logic to generate tags based on the sentence
	// Replace this with your own implementation
	return []string{"tag1", "tag2", "tag3"}
}
