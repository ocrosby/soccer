package handlers

import (
	"log"
	"net/http"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("Todo: Create user")
}

func (h *UserHandler) Read(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to handle the request
	userID := r.PathValue("userID")

	if userID == "" {
		// The userID field was not specified so return all users.
		write, err := w.Write([]byte("Hello, users"))
		if err != nil {
			log.Printf("wrote %d bytes", write)
			log.Printf("error writing response: %v", err)
		} else {
			log.Println("wrote %d bytes", write)
		}
	} else {
		// The userID field was specified so return the user with that ID.
		write, err := w.Write([]byte("Hello, " + userID))
		if err != nil {
			log.Printf("wrote %d bytes", write)
			log.Printf("error writing response: %v", err)
		} else {
			log.Println("wrote %d bytes", write)
		}
	}

	// https://pkg.go.dev/net/http#ServeMux
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Todo: Update user")
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("Todo: Delete user")
}
