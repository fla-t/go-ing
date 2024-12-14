package app

import (
	"database/sql"

	userAPI "github.com/fla-t/go-ing/UserService/api/user"
	"github.com/fla-t/go-ing/UserService/internal/services/user"
	uowInmemory "github.com/fla-t/go-ing/UserService/internal/uow/inmemory"
	uowSQL "github.com/fla-t/go-ing/UserService/internal/uow/sql"

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
