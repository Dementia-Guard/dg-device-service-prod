package repositories

import (
	"api/config"
	"api/models"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetCollection retrieves the correct collection based on the name
func GetCollection(collectionName string) *mongo.Collection {
	return config.GetCollection(collectionName)
}

// GetAllSensorData fetches sensor data based on collection name
func GetAllSensorData(collectionName string) ([]models.SensorData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data []models.SensorData
	collection := GetCollection(collectionName)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("database error while fetching sensor data: " + err.Error())
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var entry models.SensorData
		if err := cursor.Decode(&entry); err != nil {
			log.Println("Decoding error:", err)
			return nil, errors.New("failed to decode sensor data: " + err.Error())
		}
		data = append(data, entry)
	}

	if len(data) == 0 {
		return nil, errors.New("no sensor data found")
	}

	return data, nil
}

// GetSensorDataByDate fetches sensor data based on collection name and specific date
func GetSensorDataByDate(collectionName string, date time.Time) ([]models.SensorData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data []models.SensorData
	collection := GetCollection(collectionName)

	// Format the date to match the date format in MongoDB (assumes "date" field is stored in UTC)
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.Add(24 * time.Hour)

	// Query for records where the _id.date is within the specified day
	filter := bson.M{
		"_id.date": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.New("database error while fetching sensor data by date: " + err.Error())
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var entry models.SensorData
		if err := cursor.Decode(&entry); err != nil {
			log.Println("Decoding error:", err)
			return nil, errors.New("failed to decode sensor data: " + err.Error())
		}
		data = append(data, entry)
	}

	if len(data) == 0 {
		return nil, errors.New("no sensor data found for the specified date")
	}

	return data, nil
}
