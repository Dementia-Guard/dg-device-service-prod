package main

import (
	"fmt"
	"api/config"
	"api/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load ENV Variables
	config.LoadEnv()

	// Initialize Firebase
	config.InitFirebase() 

	// Initialize database first
	config.ConnectDB()

	// Get PORT from environment variable
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// Setup routes
	router := routes.SetupRouter()

	// Start server on the specified port
	fmt.Printf("🚀 Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
