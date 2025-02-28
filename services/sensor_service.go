package services

import (
	"api/models"
	"api/repositories"
	"errors"
	"time"
)

// GetSensorDataService retrieves all sensor records from the specified collection
func GetSensorDataService(collectionName string) ([]models.SensorData, error) {
	data, err := repositories.GetAllSensorData(collectionName)
	if err != nil {
		return nil, errors.New("could not retrieve sensor data: " + err.Error())
	}
	return data, nil
}

// GetSensorDataByDateService retrieves sensor records for a specific date
func GetSensorDataByDateService(collectionName string, date time.Time) ([]models.SensorData, error) {
	data, err := repositories.GetSensorDataByDate(collectionName, date)
	if err != nil {
		return nil, errors.New("could not retrieve sensor data: " + err.Error())
	}
	return data, nil
}