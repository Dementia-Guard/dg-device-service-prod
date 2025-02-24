package config

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client

func InitFirebase() {
	base64Key := os.Getenv("FIRE_CRED_KEY")
	if base64Key == "" {
		log.Fatal("❌ Firebase key not found in environment variables")
	}

	// Decode Base64 string
	decodedKey, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		log.Fatalf("❌Failed to decode Firebase key: %v", err)
	}

	// Initialize Firebase using the decoded JSON
	opt := option.WithCredentialsJSON(decodedKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("❌ Error initializing Firebase app: %v", err)
	}

	// Initialize Firestore
	FirestoreClient, err = app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Error initializing Firestore: %v", err)
	}
	// defer FirestoreClient.Close()

	log.Println("🔥✅ Firebase initialized successfully")
}

// Close Firestore when the app exits
func CloseFirestore() {
	if FirestoreClient != nil {
		err := FirestoreClient.Close()
		if err != nil {
			log.Printf("⚠️ Error closing Firestore: %v", err)
		} else {
			log.Println("✅ Firestore connection closed successfully")
		}
	}
}