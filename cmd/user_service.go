package main

import (
	"log"
	"os"
	"strconv"

	app "github.com/fla-t/go-ing/internal/app/user"
)

func main() {
	// Default values
	port := 50052
	useInMemory := true

	// Parse command-line arguments
	if len(os.Args) >= 2 {
		if p, err := strconv.Atoi(os.Args[1]); err == nil {
			port = p
		}
	}
	if len(os.Args) >= 3 {
		if b, err := strconv.ParseBool(os.Args[2]); err == nil {
			useInMemory = b
		}
	}

	// Log parsed values
	log.Printf("Starting gRPC server on port %d with InMemory=%v...\n", port, useInMemory)

	// Start the gRPC application
	app.StartGRPCApp(port, useInMemory)
}
