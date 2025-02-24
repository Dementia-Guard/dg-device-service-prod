package repositories

import (
	"api/config"
	"api/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Lazy initialization of user collection
var userCollection *mongo.Collection

func getUserCollection() *mongo.Collection {
	if userCollection == nil {
		userCollection = config.GetCollection("users")
	}
	return userCollection
}

// Fetch all users
func GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User
	cursor, err := getUserCollection().Find(ctx, bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no users found")
		}
		return nil, errors.New("database error while fetching users")
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, errors.New("failed to decode user data")
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}
