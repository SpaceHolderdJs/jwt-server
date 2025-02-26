package config

import (
	"log"
	"os"
)

var JWTSecret []byte

var MongoConnectionLink string

func init() {

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Println("WARNING: Using default JWT secret. Set JWT_SECRET environment variable in production.")
		jwtSecret = ""
	}
	JWTSecret = []byte(jwtSecret)

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Println("WARNING: Using default MongoDB connection. Set MONGO_URI environment variable in production.")
		mongoURI = ""
	}
	MongoConnectionLink = mongoURI
}
