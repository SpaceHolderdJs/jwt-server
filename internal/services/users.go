package services

import (
	// "jwt-auth-service/internal/service"
	config "jwt-auth-service"
	"jwt-auth-service/internal/models"

	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var users = []models.User{{ID: 1, Username: "SpaceHolder", Password: "12345"}}

func Login(user models.User) (string, models.User, error) {
	// Request with mongo

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.ID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(config.JWTSecret)

	if err != nil {
		return "", user, err
	}

	return tokenString, user, nil
}

func GetAllUsers() []models.User {
	return users
}

func DeleteUser(id uint) {
	users = slices.DeleteFunc(users, func(user models.User) bool {
		return user.ID != id
	})
}
