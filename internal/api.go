package internal

import (
	"github.com/ocrosby/soccer/internal/middleware"
	user_service "github.com/ocrosby/soccer/internal/user-service/handlers"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	authHandler := user_service.NewAuthorizationHandler()
	userHandler := user_service.NewUserHandler()
	tokenHandler := user_service.NewTokenHandler()
	kubernetesHandler := user_service.NewKubernetesHandler()

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

	// Wrap the entire router with the LoggingMiddleware
	wrappedRouter := middleware.LoggingMiddleware(router)

	server := http.Server{
		Addr:    s.addr,
		Handler: wrappedRouter,
	}

	log.Printf("Starting server on %s", s.addr)
	return server.ListenAndServe()
}
