package tds_coaching_change_service

import (
	"github.com/ocrosby/soccer/internal/middleware"
	httpHandlers "github.com/ocrosby/soccer/internal/tds-coaching-change-service/handlers"
	"net/http"
)

func SetupRoutes(app *Application) *http.ServeMux {
	router := http.NewServeMux()

	changesHandler := httpHandlers.NewChangesHandler()
	kubernetesHandler := httpHandlers.NewKubernetesHandler()
	swaggerUIHandler := httpHandlers.NewSwaggerUIHandler("/app/swagger-ui", "/app/swagger-ui/index.html")

	router.Handle("GET /healthz", middleware.LoggingMiddleware(http.HandlerFunc(kubernetesHandler.HealthCheckHandler)))
	router.Handle("GET /ready", middleware.LoggingMiddleware(http.HandlerFunc(kubernetesHandler.ReadinessCheckHandler)))
	router.Handle("GET /start", middleware.LoggingMiddleware(http.HandlerFunc(kubernetesHandler.StartupCheckHandler)))

	router.Handle("GET /changes", middleware.LoggingMiddleware(middleware.Cache(http.HandlerFunc(changesHandler.Read))))

	// Serve Swagger documentation
	// router.HandleFunc("/swagger.yaml", swaggerHandler.ServeHTTP)

	router.Handle("/docs/", middleware.LoggingMiddleware(http.StripPrefix("/docs/", http.HandlerFunc(swaggerUIHandler.ServeHTTP))))

	return router
}
