package routes

import (
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes all routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Global Error Recovery Middleware (Handles Panics)
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			utils.ErrorResponse(c, http.StatusInternalServerError, err, false)
		} else {
			utils.ErrorResponse(c, http.StatusInternalServerError, "An unexpected error occurred", false)
		}
	}))

	// Register all route groups
	api := router.Group("/api")
	{
		PatientRoutes(api)
		SensorRoutes(api)
	}

	// Handle 404 Not Found Routes
	router.NoRoute(func(c *gin.Context) {
		utils.ErrorResponse(c, http.StatusNotFound, "Route not found", false)
	})

	return router
}
