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
	config.LoadEnv()

	config.InitFirebase() // Initialize Firebase

	// 1️⃣ Initialize database first
	config.ConnectDB()

	// 2️⃣ Get PORT from environment variable
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// 3️⃣ Setup routes
	router := routes.SetupRouter()

	// 4️⃣ Start server on the specified port
	fmt.Printf("🚀 Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
