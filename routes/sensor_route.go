package routes

import (
	"api/controllers"
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SensorRoutes defines routes for sensor data retrieval
func SensorRoutes(api *gin.RouterGroup) {
	sensorGroup := api.Group("/sensor/data")
	{
		sensorGroup.GET("/", func(c *gin.Context) {
			utils.SuccessResponse(c, http.StatusOK, "Sensor API Online", "SUCCESS", map[string]string{"message": "Hello from Sensor API"})
		})
		// sensorGroup.GET("/data", controllers.GetSensorData)
		// Define routes for sensor data retrieval
		sensorGroup.GET("/hourly", controllers.GetHourlySensorData)
		sensorGroup.GET("/daily", controllers.GetDailySensorData)
		sensorGroup.GET("/weekly", controllers.GetWeeklySensorData)
		sensorGroup.GET("/monthly", controllers.GetMonthlySensorData)
		sensorGroup.GET("/:date", controllers.GetSpecificDateSensorData)
	}
}
