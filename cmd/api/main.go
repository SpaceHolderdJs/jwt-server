package main

import (
	"fmt"
	"jwt-auth-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", handlers.GetAll)

	fmt.Println("Server is running on port 7777")
	router.Run(":7777")
}
