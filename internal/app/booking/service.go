package user

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	bookingGRPC "github.com/fla-t/go-ing/internal/grpc/booking"
	"github.com/fla-t/go-ing/internal/services/booking"
	uowInmemory "github.com/fla-t/go-ing/internal/uow/inmemory"
	uowSQL "github.com/fla-t/go-ing/internal/uow/sql"
	proto "github.com/fla-t/go-ing/proto/booking"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGRPCApp initializes and starts the gRPC server on the specified port
func StartGRPCApp(port int, useInMemory bool) {
	var service *booking.Service

	// Initialize Unit of Work (UoW) and Service
	if useInMemory {
		service = booking.NewService(uowInmemory.NewFakeUnitOfWork())
	} else {
		db := setupDatabase()
		service = booking.NewService(uowSQL.NewDbUnitOfWork(db))
	}

	// Start the gRPC server
	startGRPCServer(service, port)
}

// setupDatabase initializes the SQLite database
func setupDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	// Create booking table if it doesn't exist
	_, err = db.Exec("create table if not exists bookings (id uuid primary key, user_id uuid not null, ride_id uuid not null, time timestamptz not null);")
	if err != nil {
		panic(err)
	}

	// Create booking table if it doesn't exist
	_, err = db.Exec("create table if not exists rides (id uuid primary key, source text not null, destination text not null, distance double not null, cost double not null, time timestamptz not null);")
	if err != nil {
		panic(err)
	}
	return db
}

// startGRPCServer starts and runs the gRPC server on the specified port
func startGRPCServer(service *booking.Service, port int) {
	address := fmt.Sprintf(":%d", port)

	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	}

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	grpcService := bookingGRPC.NewBookingService(service)

	// Register the service
	proto.RegisterBookingServiceServer(grpcServer, grpcService)
	reflection.Register(grpcServer) // Enable reflection for tools like grpcurl

	log.Printf("gRPC server started at %s\n", address)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
