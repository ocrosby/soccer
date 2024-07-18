package user_service

import (
	httpHandlers "github.com/ocrosby/soccer/internal/user-service/handlers"
	"net/http"
)

func SetupRoutes(app *Application) *http.ServeMux {
	router := http.NewServeMux()

	authHandler := httpHandlers.NewAuthorizationHandler()
	userHandler := httpHandlers.NewUserHandler()
	tokenHandler := httpHandlers.NewTokenHandler()
	kubernetesHandler := httpHandlers.NewKubernetesHandler()

	router.HandleFunc("GET /healthz", kubernetesHandler.HealthCheckHandler)
	router.HandleFunc("GET /ready", kubernetesHandler.ReadinessCheckHandler)
	router.HandleFunc("GET /start", kubernetesHandler.StartupCheckHandler)

	router.HandleFunc("GET /authorize", authHandler.HandleAuthorization)
	router.HandleFunc("POST /token", tokenHandler.HandleToken)
	router.HandleFunc("POST /revoke", tokenHandler.RevokeToken)

	router.HandleFunc("GET /users", userHandler.Read)
	router.HandleFunc("GET /users/{userID}", userHandler.Read)
	router.HandleFunc("POST /users", userHandler.Create)
	router.HandleFunc("PUT /users/{userID}", userHandler.Update)
	router.HandleFunc("DELETE /users/{userID}", userHandler.Delete)

	return router

}
