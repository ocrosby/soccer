package location_service

import (
	httpHandlers "github.com/ocrosby/soccer/internal/location-service/handlers"
	"net/http"
)

func SetupRoutes(app *Application) *http.ServeMux {
	router := http.NewServeMux()

	countriesHandler := httpHandlers.NewCountryHandler()

	router.HandleFunc("GET /countries", countriesHandler.Read)
	router.HandleFunc("GET /countries/{id}", countriesHandler.Read)
	router.HandleFunc("POST /countries", countriesHandler.Create)
	router.HandleFunc("PUT /countries/{id}", countriesHandler.Update)
	router.HandleFunc("DELETE /countries/{id}", countriesHandler.Delete)

	return router
}
