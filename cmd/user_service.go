package main

import (
	"log"

	"github.com/caarlos0/env"
	app "github.com/fla-t/go-ing/internal/app/user"
)

// Config holds the application configuration loaded from environment variables
type Config struct {
	Port        int  `env:"GRPC_PORT" envDefault:"50052"`   // gRPC port
	UseInMemory bool `env:"USE_INMEMORY" envDefault:"true"` // Use in-memory database
}

func main() {
	// Load environment variables into the Config struct
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}

	// Log the configuration
	log.Printf("Starting gRPC server on port %d with InMemory=%v...\n", cfg.Port, cfg.UseInMemory)

	// Start the gRPC application
	app.StartGRPCApp(cfg.Port, cfg.UseInMemory)
}
