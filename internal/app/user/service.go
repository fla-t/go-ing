package user

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	userGRPC "github.com/fla-t/go-ing/internal/grpc/user"
	"github.com/fla-t/go-ing/internal/services/user"
	uowInmemory "github.com/fla-t/go-ing/internal/uow/inmemory"
	uowSQL "github.com/fla-t/go-ing/internal/uow/sql"
	proto "github.com/fla-t/go-ing/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGRPCApp initializes and starts the gRPC server on the specified port
func StartGRPCApp(port int, useInMemory bool) {
	var service *user.Service

	// Initialize Unit of Work (UoW) and Service
	if useInMemory {
		service = user.NewService(uowInmemory.NewFakeUnitOfWork())
	} else {
		db := setupDatabase()
		service = user.NewService(uowSQL.NewDbUnitOfWork(db))
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

	// Create users table if it doesn't exist
	_, err = db.Exec("create table if not exists users (id uuid primary key, name text, email text)")
	if err != nil {
		panic(err)
	}
	return db
}

// startGRPCServer starts and runs the gRPC server on the specified port
func startGRPCServer(service *user.Service, port int) {
	address := fmt.Sprintf(":%d", port)

	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	}

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	grpcService := userGRPC.NewUserService(service)

	// Register the service
	proto.RegisterUserServiceServer(grpcServer, grpcService)
	reflection.Register(grpcServer) // Enable reflection for tools like grpcurl

	log.Printf("gRPC server started at %s\n", address)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
