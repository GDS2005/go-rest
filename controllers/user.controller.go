package controllers

import (
	"fmt"
	"net/http"
	"strconv" // Import strconv for type conversion

	"example/database/postgres"
	"example/models"
	"example/repositories"
	"example/services"

	"github.com/gin-gonic/gin"
)

// GetUser retrieves a user by ID
func GetUser(c *gin.Context) {
	// Get user ID from request parameters
	userIDStr := c.Param("id")

	// Convert userIDStr to integer
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Initialize database connection using the InitDB function from the postgres package
	db, err := postgres.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	// Inject user repository with the database connection
	userRepository := repositories.NewUserRepository(db)

	// Inject user service with the user repository
	userService := services.NewUserService(userRepository)

	// Get user by ID from service
	user, err := userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	// Parse request body to get user data
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Initialize database connection using the InitDB function from the postgres package
	db, err := postgres.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	// Inject user repository with the database connection
	userRepository := repositories.NewUserRepository(db)

	// Inject user service with the user repository
	userService := services.NewUserService(userRepository)

	// Create user using service
	if err := userService.CreateUser(&newUser); err != nil {
		fmt.Println("Error creating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return response
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// GetUsers retrieves all users
func GetUsers(c *gin.Context) {
	// Initialize database connection using the InitDB function from the postgres package
	db, err := postgres.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	// Inject user repository with the database connection
	userRepository := repositories.NewUserRepository(db)

	// Inject user service with the user repository
	userService := services.NewUserService(userRepository)

	// Get users from service
	users, err := userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
	// Get user ID from request parameters
	userIDStr := c.Param("id")

	// Convert userIDStr to integer
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse request body to get updated user data
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Initialize database connection using the InitDB function from the postgres package
	db, err := postgres.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	// Inject user repository with the database connection
	userRepository := repositories.NewUserRepository(db)

	// Inject user service with the user repository
	userService := services.NewUserService(userRepository)

	// Update user by ID using service
	if err := userService.UpdateUser(userID, &updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	// Get user ID from request parameters
	userIDStr := c.Param("id")

	// Convert userIDStr to integer
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Initialize database connection using the InitDB function from the postgres package
	db, err := postgres.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	// Inject user repository with the database connection
	userRepository := repositories.NewUserRepository(db)

	// Inject user service with the user repository
	userService := services.NewUserService(userRepository)

	// Delete user by ID using service
	if err := userService.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
