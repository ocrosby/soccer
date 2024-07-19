package handlers

import (
	"log"
	"net/http"
)

type CountryHandler struct {
}

func NewCountryHandler() *CountryHandler {
	return &CountryHandler{}
}

func (h *CountryHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("Todo: Create country")
}

func (h *CountryHandler) Read(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to handle the request
	countryID := r.PathValue("countryID")

	if countryID == "" {
		// The countryID field was not specified so return all countries.
		getAllCountries(w)
	} else {
		// The countryID field was specified so return the country with that ID.
		getCountryByID(w, countryID)
	}

	// https://pkg.go.dev/net/http#ServeMux
}

func getCountryByID(w http.ResponseWriter, countryID string) {
	write, err := w.Write([]byte("Hello, " + countryID))
	if err != nil {
		log.Printf("wrote %d bytes", write)
		log.Printf("error writing response: %v", err)
	} else {
		log.Println("wrote %d bytes", write)
	}
}

func getAllCountries(w http.ResponseWriter) {
	write, err := w.Write([]byte("Hello, countries"))
	if err != nil {
		log.Printf("wrote %d bytes", write)
		log.Printf("error writing response: %v", err)
	} else {
		log.Println("wrote %d bytes", write)
	}
}

func (h *CountryHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Todo: Update country")
}

func (h *CountryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("Todo: Delete country")
}
