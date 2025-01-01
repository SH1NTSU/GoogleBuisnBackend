package db

import (
    "context"
    "fmt"
    "log"
    "sync"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// Singleton pattern
var (
    clientInstance *mongo.Client
    clientOnce     sync.Once
)

// InitDB initializes and returns the MongoDB client instance
func InitDB() *mongo.Client {
    clientOnce.Do(func() {
        // MongoDB connection URI
        clientOptions := options.Client().ApplyURI("mongodb+srv://GoogleBuisn:GoogleBuisn2024@cluster0.kkfzp.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

        // Connect to MongoDB
        var err error
        clientInstance, err = mongo.Connect(context.Background(), clientOptions)
        if err != nil {
            log.Fatalf("Failed to connect to MongoDB: %v", err)
        }

        // Ensure connection is established
        err = clientInstance.Ping(context.Background(), nil)
        if err != nil {
            log.Fatalf("Failed to ping MongoDB: %v", err)
        }

        fmt.Println("Successfully connected to MongoDB!")
    })

    return clientInstance
}

// GetClient returns the MongoDB client instance
func GetClient() *mongo.Client {
    if clientInstance == nil {
        log.Fatal("Database not initialized. Call InitDB first.")
    }
    return clientInstance
}

