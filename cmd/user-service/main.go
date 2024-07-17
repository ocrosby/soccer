package main

import (
	"github.com/ocrosby/soccer/internal"
	"log"
)

func main() {
	server := internal.NewAPIServer(":8080")

	if err := server.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
