package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjk1995/codegen-oauthapi/api"
)

func main() {
	router := gin.Default()

	// Load the OpenAPI spec
	apiSpec, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("failed to load OpenAPI spec: %v", err)
	}
	apiSpec.Servers = nil // Ensure no server URL in spec

	// Register routes
	api.RegisterHandlers(router, &Server{})

	// Start server
	router.Run(":8080")
}

// Server struct implementing the generated interface
type Server struct{}

// CreateUser implements api.ServerInterface.
func (s *Server) CreateUser(c *gin.Context) {
	var newUser struct {
		Id    *string `json:"id"`
		Name  string  `json:"name"`
		Email string  `json:"email"`
	}
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	newUser.Id = new(string)
	*newUser.Id = "123"
	c.JSON(http.StatusCreated, newUser)
}

// DeleteUser implements api.ServerInterface.
func (s *Server) DeleteUser(c *gin.Context, userId string) {
	// Simulate user deletion
	c.JSON(http.StatusOK, gin.H{"status": "User deleted", "userId": userId})
}

// GetUserById implements api.ServerInterface.
func (s *Server) GetUserById(c *gin.Context, userId string) {
	user := struct {
		Id    *string `json:"id"`
		Name  string  `json:"name"`
		Email string  `json:"email"`
	}{
		Id:    &userId,
		Name:  "John Doe",
		Email: "john@example.com",
	}
	c.JSON(http.StatusOK, user)
}

// GetUsers implements api.ServerInterface.
func (s *Server) GetUsers(c *gin.Context) {
	users := []struct {
		Id    *string `json:"id"`
		Name  string  `json:"name"`
		Email string  `json:"email"`
	}{
		{Id: new(string), Name: "John Doe", Email: "john@example.com"},
		{Id: new(string), Name: "Jane Smith", Email: "jane@example.com"},
	}
	*users[0].Id = "1"
	*users[1].Id = "2"
	c.JSON(http.StatusOK, users)
}

// PatchUser implements api.ServerInterface.
func (s *Server) PatchUser(c *gin.Context, userId string) {
	var updates struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.BindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	updatedUser := struct {
		Id    *string `json:"id"`
		Name  string  `json:"name"`
		Email string  `json:"email"`
	}{
		Id:    &userId,
		Name:  updates.Name,
		Email: updates.Email,
	}
	c.JSON(http.StatusOK, updatedUser)
}

// UpdateUser implements api.ServerInterface.
func (s *Server) UpdateUser(c *gin.Context, userId string) {
	var updatedUser struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// Simulate updating a user
	updatedUser.Id = userId
	c.JSON(http.StatusOK, updatedUser)
}

// Implement the GET /hello endpoint
func (s *Server) GetHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello from Gin-OAPI!"})
}
