package handlers

import (
	"context"
	"encoding/json"
	"github.com/ocrosby/soccer/internal/location-service/database/models"
	"github.com/ocrosby/soccer/internal/location-service/database/repository"
	"log"
	"net/http"
)

type CountryHandler struct {
	repo repository.CountryRepositoryInterface
}

func NewCountryHandler(repo repository.CountryRepositoryInterface) *CountryHandler {
	return &CountryHandler{
		repo: repo,
	}
}

func (h *CountryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var country models.Country

	if err := json.NewDecoder(r.Body).Decode(&country); err != nil {
		log.Printf("Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(context.Background(), country); err != nil {
		log.Printf("Error creating country: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(country)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CountryHandler) Read(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to handle the request
	countryID := r.PathValue("id")

	if countryID == "" {
		// The countryID field was not specified so return all countries.
		getAllCountries(h.repo, w)
	} else {
		// The countryID field was specified so return the country with that ID.
		getCountryByID(h.repo, w, countryID)
	}

	// https://pkg.go.dev/net/http#ServeMux
}

func getCountryByID(repo repository.CountryRepositoryInterface, w http.ResponseWriter, id string) {
	country, err := repo.FindById(context.Background(), id)
	if err != nil {
		log.Printf("Error finding country: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if country == nil {
		log.Printf("Country not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(country)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getAllCountries(repo repository.CountryRepositoryInterface, w http.ResponseWriter) {
	countries, err := repo.FindAll(context.Background())
	if err != nil {
		log.Printf("Error finding countries: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(countries)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *CountryHandler) Update(w http.ResponseWriter, r *http.Request) {
	var country models.Country

	if err := json.NewDecoder(r.Body).Decode(&country); err != nil {
		log.Printf("Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(context.Background(), country); err != nil {
		log.Printf("Error updating country: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CountryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		log.Printf("Country ID not provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(context.Background(), id); err != nil {
		log.Printf("Error deleting country: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
