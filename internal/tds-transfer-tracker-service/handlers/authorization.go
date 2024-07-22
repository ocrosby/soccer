package handlers

import (
	"log"
	"net/http"
)

type AuthorizationHandler struct {
}

func NewAuthorizationHandler() *AuthorizationHandler {
	return &AuthorizationHandler{}
}

func (h *AuthorizationHandler) HandleAuthorization(w http.ResponseWriter, r *http.Request) {
	// Implement the authorization logic here
	log.Printf("Todo: Handle authorization")
}
