package controllers

import (
	"api/services"
	"api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetHourlySensorData retrieves all hourly sensor records
func GetHourlySensorData(c *gin.Context) {
	data, err := services.GetSensorDataService("sensor_data_hourly")
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Hourly Sensor Data Retrieved", true, data)
}

// GetDailySensorData retrieves all daily sensor records
func GetDailySensorData(c *gin.Context) {
	data, err := services.GetSensorDataService("sensor_data_daily")
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Daily Sensor Data Retrieved", true, data)
}

// GetWeeklySensorData retrieves all weekly sensor records
func GetWeeklySensorData(c *gin.Context) {
	data, err := services.GetSensorDataService("sensor_data_weekly")
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Weekly Sensor Data Retrieved", true, data)
}

// GetMonthlySensorData retrieves all monthly sensor records
func GetMonthlySensorData(c *gin.Context) {
	data, err := services.GetSensorDataService("sensor_data_monthly")
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Monthly Sensor Data Retrieved", true, data)
}

// GetSpecificDateSensorData retrieves sensor records for a specific date
func GetSpecificDateSensorData(c *gin.Context) {
	// Parse the date parameter from the request query string
	dateStr := c.Param("date")
	if dateStr == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Date parameter is required",false)
		return
	}

	// Parse the date string to a DateTime object
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid date format, use YYYY-MM-DD", false)
		return
	}

	// Call the service to retrieve sensor data for the specific date
	data, err := services.GetSensorDataByDateService("sensor_data_daily", parsedDate)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sensor Data Retrieved for Specific Date", true, data)
}
