package user

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	acl "github.com/fla-t/go-ing/internal/acl/user/grpc"
	bookingGRPC "github.com/fla-t/go-ing/internal/grpc/booking"
	"github.com/fla-t/go-ing/internal/services/booking"
	uowInmemory "github.com/fla-t/go-ing/internal/uow/inmemory"
	uowSQL "github.com/fla-t/go-ing/internal/uow/sql"
	proto "github.com/fla-t/go-ing/proto/booking"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

// StartGRPCApp initializes and starts the gRPC server on the specified port
func StartGRPCApp(port int, useInMemory bool, userServiceAddress string) {
	var service *booking.Service

	// Setup gRPC connection for User Service ACL
	userConn := setupUserGRPCConnection(userServiceAddress)
	// defer userConn.Close()

	userACL := acl.NewGRPCUserACL(userConn)

	// Initialize Unit of Work (UoW) and Service
	if useInMemory {
		service = booking.NewService(uowInmemory.NewFakeUnitOfWork(), userACL)
	} else {
		db := setupDatabase()
		defer db.Close() // close the db connection

		service = booking.NewService(uowSQL.NewDbUnitOfWork(db), userACL)
	}

	// Start the gRPC server
	startGRPCServer(service, port)

}

// setupDatabase initializes the SQLite database
func setupDatabase() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Create booking table if it doesn't exist
	_, err = db.Exec("create table if not exists bookings (id uuid primary key, user_id uuid not null, ride_id uuid not null, time timestamptz not null);")
	if err != nil {
		panic(err)
	}

	// Create booking table if it doesn't exist
	_, err = db.Exec("create table if not exists rides (id uuid primary key, source text not null, destination text not null, distance double precision not null, cost double precision not null);")
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

// setupUserGRPCConnection establishes a gRPC connection to the User Service
func setupUserGRPCConnection(userServiceAddress string) *grpc.ClientConn {
	conn, err := grpc.Dial(userServiceAddress, grpc.WithInsecure(), grpc.WithIdleTimeout(0))
	if err != nil {
		log.Fatalf("Failed to connect to User Service at %s: %v", userServiceAddress, err)
	}
	log.Printf("Connected to User Service at %s\n", userServiceAddress)
	return conn
}
