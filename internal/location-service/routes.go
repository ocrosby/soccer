package location_service

import (
	"github.com/ocrosby/soccer/internal/location-service/database/repository"
	httpHandlers "github.com/ocrosby/soccer/internal/location-service/handlers"
	"net/http"
)

func SetupRoutes(app *Application) *http.ServeMux {
	router := http.NewServeMux()

	repo := repository.NewCountryRepository(app.DB)
	countriesHandler := httpHandlers.NewCountryHandler(repo)
	kubernetesHandler := httpHandlers.NewKubernetesHandler(app.DB)

	router.HandleFunc("GET /healthz", kubernetesHandler.HealthCheckHandler)
	router.HandleFunc("GET /ready", kubernetesHandler.ReadinessCheckHandler)
	router.HandleFunc("GET /start", kubernetesHandler.StartupCheckHandler)

	router.HandleFunc("GET /countries", countriesHandler.Read)
	router.HandleFunc("GET /countries/{id}", countriesHandler.Read)
	router.HandleFunc("POST /countries", countriesHandler.Create)
	router.HandleFunc("PUT /countries/{id}", countriesHandler.Update)
	router.HandleFunc("DELETE /countries/{id}", countriesHandler.Delete)

	return router
}
