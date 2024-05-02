package main

import (
	"example/controllers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Allow any origin (for testing purposes)
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"}
	config.ExposeHeaders = []string{"Authorization"}

	r.Use(cors.New(config))

	// Define your API routes
	r.GET("/v1/users", controllers.GetUsers)
	r.GET("/v1/users/:id", controllers.GetUser)
	r.POST("/v1/users", controllers.CreateUser)
	r.PATCH("/v1/users/:id", controllers.UpdateUser)
	r.DELETE("/v1/users/:id", controllers.DeleteUser)

	// Start the HTTP server
	port := os.Getenv("PORT") // Set your desired port
	if port == "" {
		port = "8080" // Default port if not specified
	}
	r.Run(":" + port)
}
