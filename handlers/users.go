package handlers

import (
	"fmt"
	"jwt-auth-service/internal/models"
	"jwt-auth-service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary     Login user
// @Description Authenticate a user and return a JWT token
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       user body models.User true "Login credentials"
// @Success     200 {object} object{token=string,user=models.User} "Successfully authenticated"
// @Failure     400 {object} object{error=string} "Invalid request or credentials"
// @Router      /auth/login [post]
func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, user, err := services.Login(user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		fmt.Println(err, "err")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

// @Summary     Register user
// @Description Register a new user and return a JWT token
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       user body models.User true "User registration details"
// @Success     201 {object} object{token=string,user=string} "Successfully registered"
// @Failure     400 {object} object{error=string} "Invalid request"
// @Router      /auth/register [post]
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, id, err := services.Register(services.RegisterData{
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err, "err")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token, "userId": id})
}

// @Summary     Get all users
// @Description Retrieve all users
// @Tags        users
// @Produce     json
// @Success     200 {array} models.User "List of users"
// @Failure     500 {object} object{error=string} "Server error"
// @Router      /api/users [get]
// @Security    Bearer
func GetAll(c *gin.Context) {
	users, err := services.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary     Delete user
// @Description Delete a user by ID
// @Tags        users
// @Produce     json
// @Param       id path string true "User ID"
// @Success     200 {object} object{deleted=bool} "User deleted"
// @Failure     400 {object} object{error=string} "Invalid ID format"
// @Failure     404 {object} object{error=string} "User not found"
// @Router      /api/users/{id} [delete]
// @Security    Bearer
func Delete(c *gin.Context) {
	id := c.Param("id")

	success := services.DeleteUser(id)

	if !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
