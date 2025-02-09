package handlers

import (
	"fmt"
	"jwt-auth-service/internal/models"
	"jwt-auth-service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Validation

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}

	token, user, err := services.Login(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid token"})
		fmt.Println(err, "err")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

func GetAll(c *gin.Context) {
	c.JSON(200, services.GetAllUsers())
}
