package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Encode writes a JSON object to an HTTP response.
func Encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	log.Printf("Encoding response for %s %s, status: %d", r.Method, r.URL, status)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("Encode json: %w", err)
	}

	return nil
}

// Decode reads a JSON object from an HTTP request and stores it in v.
func Decode[T any](r *http.Request) (T, error) {
	var v T

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("Decode json: %w", err)
	}

	return v, nil
}
