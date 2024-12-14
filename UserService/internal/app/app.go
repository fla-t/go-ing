package app

import (
	"database/sql"
	"log"
	"net"

	userAPI "github.com/fla-t/go-ing/UserService/internal/api/user"
	userGRPC "github.com/fla-t/go-ing/UserService/internal/grpc/user"
	"github.com/fla-t/go-ing/UserService/internal/services/user"
	uowInmemory "github.com/fla-t/go-ing/UserService/internal/uow/inmemory"
	uowSQL "github.com/fla-t/go-ing/UserService/internal/uow/sql"
	"github.com/fla-t/go-ing/UserService/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gin-gonic/gin"
)

// App holds the router
type App struct {
	Router *gin.Engine
}

// NewApp creates a new App
func NewApp(useInMemory bool) *App {
	// Init unit of work

	var service *user.Service

	if useInMemory {
		// Init in-memory repository
		service = user.NewService(uowInmemory.NewFakeUnitOfWork())
	} else {
		// Init DB repository
		db := setupDatabase()
		service = user.NewService(uowSQL.NewDbUnitOfWork(db))
	}

	// Init handler
	userhandler := userAPI.NewUserHandler(service)

	// Setup routes
	router := gin.Default()
	setupRoutes(router, userhandler)

	// Start gRPC server
	go startGRPCServer(service)

	return &App{Router: router}
}

// SetupDatabase sets up the database
func setupDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("create table if not exists users (id text primary key, name text, email text)")
	if err != nil {
		panic(err)
	}
	return db
}

// SetupRoutes sets up the routes
func setupRoutes(router *gin.Engine, handler *userAPI.Handler) {
	r := router.Group("/users")
	{
		r.POST("/", handler.CreateUser)
		r.GET("/:id", handler.GetUserByID)
	}
}

// startGRPCServer starts the gRPC server
func startGRPCServer(service *user.Service) {
	listen, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpcService := userGRPC.NewUserService(service)

	proto.RegisterUserServiceServer(grpcServer, grpcService)
	reflection.Register(grpcServer)

	log.Println("gRPC server started at :50051")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
