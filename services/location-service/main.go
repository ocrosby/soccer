package main

import (
	"context"
	"database/sql"
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

	dbConnectionString := getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}(db)

	// Ensure the database is reachable
	if err = db.PingContext(ctx); err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}
	
	app := service.NewApplication(":8080", db)
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
