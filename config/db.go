package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database // Global database instance

// ConnectDB initializes MongoDB connection
func ConnectDB() {
	mongoURI := os.Getenv("MONGO_URI")
	mongo_db_name := os.Getenv("DB_NAME")

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("❌ Failed to create MongoDB client:", err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatal("❌ MongoDB connection failed:", err)
	}

	// Assign global MongoDB instance
	MongoDB = client.Database(mongo_db_name)
	fmt.Println("✅ Connected to MongoDB successfully!")
}

// GetCollection safely fetches a collection
func GetCollection(collectionName string) *mongo.Collection {
	if MongoDB == nil {
		log.Fatal("❌ Database not initialized. Call ConnectDB() first.")
	}
	return MongoDB.Collection(collectionName)
}
