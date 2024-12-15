package main

import (
	"log"
	"os"
	"strconv"

	app "github.com/fla-t/go-ing/internal/app/booking"
)

func main() {
	// Default values
	port := 50052
	useInMemory := true
	userServiceAddress := "localhost:50051"

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
	if len(os.Args) >= 4 {
		userServiceAddress = os.Args[3]
	}

	// Log parsed values
	log.Printf("Starting gRPC server on port %d with InMemory=%v...\n", port, useInMemory)
	log.Printf("User Service Address: %s\n", userServiceAddress)

	// Start the gRPC application
	app.StartGRPCApp(port, useInMemory, userServiceAddress)
}
