package handlers

import (
	"log"
	"net/http"
)

type ConferencesHandler struct {
}

func NewConferencesHandler() *ConferencesHandler {
	return &ConferencesHandler{}
}

func (h *ConferencesHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to handle the request
	log.Printf("Todo: Create conference")
}

func (h *ConferencesHandler) Read(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to handle the request
	id := r.PathValue("id")

	// Read the optional gender and division query parameters
	queryParams := r.URL.Query()

	gender := queryParams.Get("gender")
	if gender != "" {
		log.Printf("Gender: %s", gender)
		// Additional processing based on gender
	}

	division := queryParams.Get("division")
	if division != "" {
		log.Printf("Division: %s", division)
		// Additional processing
	}

	if id == "" {
		// The id field was not specified.
		log.Printf("Todo: Read all conferences")
	} else {
		// The id field was specified.
		log.Printf("Todo: Read conference with id %s", id)
	}
}
