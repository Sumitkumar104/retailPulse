package main

import (
    "log"
    "net/http"
    "retailPulse/models"
    "retailPulse/routes"
)

func main() {
	// Load Store Master
	err := models.LoadStoreMaster("storeMaster.csv")
	if err != nil {
		log.Fatalf("Failed to load store master: %v", err)
	}

	// Initialize Routes
	router := routes.InitializeRoutes()

	// Start Server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

