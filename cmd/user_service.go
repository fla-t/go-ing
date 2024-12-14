package main

import (
	"flag"
	"log"

	"github.com/fla-t/go-ing/internal/app"
)

func main() {
	// Command-line flags
	var port int
	var useInMemory bool

	flag.IntVar(&port, "port", 50051, "Port to run the gRPC server")
	flag.BoolVar(&useInMemory, "inmemory", true, "Use in-memory database (true/false)")
	flag.Parse()

	// Start the gRPC application
	log.Printf("Starting gRPC server on port %d with InMemory=%v...\n", port, useInMemory)
	app.StartGRPCApp(port, useInMemory)
}
