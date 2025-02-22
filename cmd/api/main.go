package main

import (
	"fmt"
	"jwt-auth-service/internal/handlers"
	"jwt-auth-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/login", handlers.Login)
		auth.POST("/register", func(ctx *gin.Context) {})
	}

	protected := router.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/users", handlers.GetAll)
	}

	router.GET("/basic", handlers.Basic)

	fmt.Println("Server is running on port 7777")
	router.Run("0.0.0.0:8080")
}
