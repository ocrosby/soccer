package main

import (
	"context"
	"fmt"
	service "github.com/ocrosby/soccer/internal/location-service"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func getEnv(variableName string) string {
	value, exists := os.LookupEnv(variableName)
	if !exists {
		// Return an empty string or a default value if the environment variable is not set
		return ""
	}

	return value
}

func run(ctx context.Context, args []string, getenv func(string) string, stdin io.Reader, stdout, stderr io.Writer) error {

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// ...

	app := service.NewApplication(":8080")
	router := service.SetupRoutes(app)

	server := http.Server{
		Addr:    app.Address,
		Handler: router,
	}

	log.Printf("Starting server on %s", app.Address)
	return server.ListenAndServe()
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, getEnv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		if _, err = fmt.Fprintf(os.Stderr, "%v\n", err); err != nil {
			os.Exit(2)
		}

		os.Exit(1)
	}
}
