package location_service

import (
	httpHandlers "github.com/ocrosby/soccer/internal/location-service/handlers"
	"net/http"
)

func SetupRoutes(app *Application) *http.ServeMux {
	router := http.NewServeMux()

	countriesHandler := httpHandlers.NewCountryHandler()

	router.HandleFunc("GET /countries", countriesHandler.Read)
	router.HandleFunc("GET /countries/{countryID}", countriesHandler.Read)
	router.HandleFunc("POST /countries", countriesHandler.Create)
	router.HandleFunc("PUT /countries/{countryID}", countriesHandler.Update)
	router.HandleFunc("DELETE /countries/{countryID}", countriesHandler.Delete)

	return router
}
