package user

import (
	"net/http"

	"github.com/fla-t/go-ing/internal/services/user"

	"github.com/gin-gonic/gin"
)

// Handler holds all the functions that can be called from the user API
type Handler struct {
	service *user.Service
}

// NewUserHandler creates a new Handler
func NewUserHandler(service *user.Service) *Handler {
	return &Handler{service: service}
}

// CreateUser handles POST /users
func (h *Handler) CreateUser(c *gin.Context) {
	var u User

	// Bind the request body to the user struct
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Validate the user
	if err := u.Validate(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Call the service to create the user
	if err := h.service.CreateUser(u.ConvertToDomain()); err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "User created successfully"})
}

// GetUserByID handles GET /users/:id
func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	// Call the service to get the user by its id
	u, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, u)
}
