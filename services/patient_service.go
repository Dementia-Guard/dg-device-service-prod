package services

import (
	"api/models"
	"api/repositories"
	"errors"
)

// Fetch all patients
func GetPatientsService() ([]models.Patient, error) {
	patients, err := repositories.GetAllPatients()
	if err != nil {
		return nil, errors.New("could not fetch patients")
	}
	return patients, nil
}

// GetPatientByIdService fetches a single patient by ID
func GetPatientByIdService(patientId string) (*models.Patient, error) {
	if patientId == "" {
		return nil, errors.New("patient ID is required")
	}
	return repositories.GetPatientById(patientId)
}

func EditPatientByIdService(patientId string, updatedPatient *models.Patient) (*models.Patient, error) {
    if patientId == "" {
        return nil, errors.New("patient ID is required")
    }
    return repositories.UpdatePatientById(patientId, updatedPatient)
}


// Create a new patient
// func CreatePatientService(patient models.Patient) (*models.Patient, error) {
// 	newPatient, err := repositories.AddPatient(patient)
// 	if err != nil {
// 		return nil, errors.New("failed to create patient")
// 	}
// 	return newPatient, nil
// }
