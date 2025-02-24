package controllers

import (
	"api/services"
	"api/utils"
	// "api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Fetch all patients
func GetPatients(c *gin.Context) {
	patients, err := services.GetPatientsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch patients"})
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Patients fetched successfully", "SUCCESS", patients)
}

// GetPatientById handles fetching a patient by ID
func GetPatientById(c *gin.Context) {
	patientId := c.Param("id")
	if patientId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient ID is required"})
		return
	}

	patient, err := services.GetPatientByIdService(patientId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Patient fetched successfully", "SUCCESS", patient)
}

// Create a new patient
// func CreatePatient(c *gin.Context) {
// 	var patient models.Patient
// 	if err := c.ShouldBindJSON(&patient); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
// 		return
// 	}

// 	newPatient, err := services.CreatePatientService(patient)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, newPatient)
// }