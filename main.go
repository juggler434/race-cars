package main

import (
	"log"
	"net/http"
	"os"

	"race-cars/internal/config"
	"race-cars/internal/middleware"
	"race-cars/internal/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Load environment variables
	if err := config.LoadEnv(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Initialize database
	if err := config.InitDB(); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	// Create router
	router := mux.NewRouter()

	// Add middleware
	router.Use(middleware.CORS)
	router.Use(middleware.Logger)
	router.Use(middleware.Recovery)

	// Setup routes
	routes.SetupRoutes(router)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Error starting server:", err)
	}
} 