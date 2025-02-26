package mongo

import (
	"context"
	"fmt"
	config "jwt-auth-service"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	// Client is the MongoDB client instance
	Client *mongo.Client
	// DB is the database instance
	DB *mongo.Database
)

// Connect establishes a connection to MongoDB
func Connect(uri string, dbName string) error {
	if uri == "" {
		uri = string(config.MongoConnectionLink)
	}

	if dbName == "" {
		dbName = "auth_service" // Default database name
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the MongoDB server to verify connection
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	// Set the database
	DB = Client.Database(dbName)

	fmt.Println("Connected to MongoDB!")
	return nil
}

// Disconnect closes the MongoDB connection
func Disconnect() error {
	if Client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := Client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %v", err)
	}

	fmt.Println("Disconnected from MongoDB")
	return nil
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
