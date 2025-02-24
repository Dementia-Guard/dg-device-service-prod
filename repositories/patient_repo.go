package repositories

import (
	"api/config"
	"api/models"
	"context"
	"errors"
	"log"

	"google.golang.org/api/iterator"
)

// Lazy initialization of patient collection
var patientCollection string = "patients"

// Fetch all patients from Firestore
func GetAllPatients() ([]models.Patient, error) {
	ctx := context.Background()

	// Initialize an empty slice to hold patient data
	var patients []models.Patient

	// Fetch documents from the "patients" collection
	iter := config.FirestoreClient.Collection(patientCollection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("❌ Error while fetching patients from Firestore: %v", err)
			return nil, errors.New("error while fetching patients")
		}

		// Map Firestore document data to Patient model
		var patient models.Patient
		if err := doc.DataTo(&patient); err != nil {
			log.Printf("❌ Error while decoding patient data: %v", err)
			return nil, errors.New("failed to decode patient data")
		}

		// Append patient to the slice
		patients = append(patients, patient)
	}

	// If no patients were found
	if len(patients) == 0 {
		return nil, errors.New("no patients found")
	}

	return patients, nil
}

// Fetch a single patient by their ID
func GetPatientById(patientId string) (*models.Patient, error) {
	ctx := context.Background()

	// Fetch the document for the given patient ID
	doc, err := config.FirestoreClient.Collection(patientCollection).Doc(patientId).Get(ctx)
	if err != nil {
		log.Printf("❌ Error fetching patient with ID %s: %v", patientId, err)
		return nil, errors.New("patient not found")
	}

	// Map Firestore document data to Patient model
	var patient models.Patient
	if err := doc.DataTo(&patient); err != nil {
		log.Printf("❌ Error while decoding patient data: %v", err)
		return nil, errors.New("failed to decode patient data")
	}

	return &patient, nil
}

// Add a new patient to Firestore
// func AddPatient(patient models.Patient) (*models.Patient, error) {
// 	ctx := context.Background()

// 	// Create a new document with auto-generated ID in Firestore
// 	ref, err := config.FirestoreClient.Collection(patientCollection).Add(ctx, patient)
// 	if err != nil {
// 		log.Printf("❌ Error adding patient to Firestore: %v", err)
// 		return nil, errors.New("failed to add patient")
// 	}

// 	// After adding the document, fetch the full document with the auto-generated ID
// 	doc, err := ref.Get(ctx)
// 	if err != nil {
// 		log.Printf("❌ Error fetching added patient data: %v", err)
// 		return nil, errors.New("failed to fetch added patient data")
// 	}

// 	var newPatient models.Patient
// 	if err := doc.DataTo(&newPatient); err != nil {
// 		log.Printf("❌ Error decoding added patient data: %v", err)
// 		return nil, errors.New("failed to decode added patient data")
// 	}

// 	return &newPatient, nil
// }

// Update a patient in Firestore
func UpdatePatient(patientId string, updatedPatient models.Patient) (*models.Patient, error) {
	ctx := context.Background()

	// Update the document with the provided patient ID
	_, err := config.FirestoreClient.Collection(patientCollection).Doc(patientId).Set(ctx, updatedPatient)
	if err != nil {
		log.Printf("❌ Error updating patient with ID %s: %v", patientId, err)
		return nil, errors.New("failed to update patient")
	}

	// Fetch the updated document
	doc, err := config.FirestoreClient.Collection(patientCollection).Doc(patientId).Get(ctx)
	if err != nil {
		log.Printf("❌ Error fetching updated patient: %v", err)
		return nil, errors.New("failed to fetch updated patient data")
	}

	var patient models.Patient
	if err := doc.DataTo(&patient); err != nil {
		log.Printf("❌ Error decoding updated patient data: %v", err)
		return nil, errors.New("failed to decode updated patient data")
	}

	return &patient, nil
}

// Delete a patient by their ID
func DeletePatient(patientId string) error {
	ctx := context.Background()

	// Delete the document with the provided patient ID
	_, err := config.FirestoreClient.Collection(patientCollection).Doc(patientId).Delete(ctx)
	if err != nil {
		log.Printf("❌ Error deleting patient with ID %s: %v", patientId, err)
		return errors.New("failed to delete patient")
	}

	return nil
}
