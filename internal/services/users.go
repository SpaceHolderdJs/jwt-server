package services

import (
	"context"
	"errors"
	config "jwt-auth-service"
	"jwt-auth-service/internal/models"
	"jwt-auth-service/internal/mongo"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const userCollection = "users"

func Login(user models.User) (string, models.User, error) {

	collection := mongo.DB.Collection(userCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var foundUser models.User
	err := collection.FindOne(ctx, bson.M{
		"email":    user.Email,
		"password": user.Password,
	}).Decode(&foundUser)

	if err != nil {
		return "", models.User{}, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": foundUser.ID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(config.JWTSecret)
	if err != nil {
		return "", models.User{}, err
	}

	return tokenString, foundUser, nil
}

type RegisterData struct {
	Email, Password string
}

func Register(data RegisterData) (string, string, error) {
	// Get users collection
	collection := mongo.DB.Collection(userCollection)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if user already exists
	count, err := collection.CountDocuments(ctx, bson.M{"email": data.Email})
	if err != nil {
		return "", "", err
	}

	if count > 0 {
		return "", "", errors.New("user already exists")
	}

	// Create new user
	newUser := models.User{
		Email:     data.Email,
		Password:  data.Password, // Note: In production, you should hash passwords
		CreatedAt: time.Now(),
	}

	// Insert user into database
	result, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return "", "", err
	}

	// Get the ID of the inserted user
	id := result.InsertedID.(primitive.ObjectID).Hex()

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": id,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(config.JWTSecret)
	if err != nil {
		return "", "", err
	}

	return tokenString, id, nil
}

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]models.User, error) {
	// Get users collection
	collection := mongo.DB.Collection(userCollection)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find all users, but don't return passwords
	opts := options.Find().SetProjection(bson.M{"password": 0})
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode users
	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id string) bool {
	// Get users collection
	collection := mongo.DB.Collection(userCollection)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}

	// Delete user
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return false
	}

	// Return true if a document was deleted
	return result.DeletedCount > 0
}
