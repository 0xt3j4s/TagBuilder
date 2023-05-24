package dto

type TagDto struct {
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	// Add other fields as needed for your tag DTO
}
