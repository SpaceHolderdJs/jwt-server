package services

import (
	// "jwt-auth-service/internal/service"
	"jwt-auth-service/internal/models"
	"slices"
)

var users = []models.User{{ID: 1, Username: "SpaceHolder", Password: "12345"}}

func GetAllUsers() []models.User {
	return users
}

func DeleteUser(id uint) {
	users = slices.DeleteFunc(users, func(user models.User) bool {
		return user.ID != id
	})
}
