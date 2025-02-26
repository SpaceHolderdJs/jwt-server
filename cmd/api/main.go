package main

import (
	"fmt"
	"jwt-auth-service/handlers"
	"jwt-auth-service/internal/middleware"
	"jwt-auth-service/internal/mongo"

	"log"

	_ "jwt-auth-service/cmd/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "jwt-auth-service/cmd/api/docs" // Add this line
)

func init() {
	// Then add this after your import statements:
	docs.SwaggerInfo.BasePath = "/"
}

// @title           JWT Auth Service API
// @version         1.0
// @description     A secure API with JWT authentication
// @host            localhost:8080
// @BasePath        /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Connect to MongoDB
	err := mongo.Connect("", "jwt-server") // Use default connection string and database name
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongo.Disconnect() // Properly close the connection when the app exits

	router := gin.Default()

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/login", handlers.Login)
		auth.POST("/register", handlers.Register)
	}

	protected := router.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/users", handlers.GetAll)
		protected.DELETE("/users/:id", handlers.Delete)
	}

	router.GET("/basic", handlers.Basic)

	fmt.Println("Server is running on port 8080")
	router.Run("0.0.0.0:8080")
}
