package routes

import (
	"api/controllers"
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PatientRoutes defines routes for patient-related operations
func PatientRoutes(api *gin.RouterGroup) {
	patientGroup := api.Group("/patient")
	{
		// Route to check if the patient API is online
		patientGroup.GET("/", func(c *gin.Context) {
			utils.SuccessResponse(c, http.StatusOK, "Patient API Online", "SUCCESS", map[string]string{"message": "Hello From Patient API"})
		})

		// Trigger error for testing purposes
		patientGroup.GET("/trigger-error", func(c *gin.Context) {
			panic("Something went wrong!")
		})

		// Route to get all patients
		patientGroup.GET("/patients", controllers.GetPatients)

		// Route to get a single patient by ID
		patientGroup.GET("/patients/:id", controllers.GetPatientById)

		// Route to create a new patient
		// patientGroup.POST("/patients", controllers.CreatePatient)
	}
}