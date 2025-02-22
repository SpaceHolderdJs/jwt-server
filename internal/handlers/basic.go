package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Basic(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello", "time": time.Now()})
}
