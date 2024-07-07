package main

import (
	"fmt"
	"net/http"
)

// HomeHandler responds to HTTP GET requests on the root "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	numberOfBytesWritten, err := fmt.Fprintln(w, "Welcome to the Gorilla Mux Service!")
	if err != nil {
		format := "Internal Server Error. Bytes written before error: %d"
		errMsg := fmt.Sprintf(format, numberOfBytesWritten)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
}

// HealthCheckHandler responds to HTTP GET requests on the "/health" endpoint
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	numberOfBytesWritten, err := fmt.Fprintln(w, "OK")
	if err != nil {
		format := "Internal Server Error. Bytes written before error: %d"
		errMsg := fmt.Sprintf(format, numberOfBytesWritten)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
}
