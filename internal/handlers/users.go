package handlers

import (
	"github.com/gin-gonic/gin"
	"jwt-auth-service/internal/services"
)

func GetAll(c *gin.Context) {
	c.JSON(200, services.GetAllUsers())
}
