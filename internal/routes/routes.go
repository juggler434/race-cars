package routes

import (
	"net/http"
	"race-cars/internal/handlers"

	"github.com/gorilla/mux"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *mux.Router) {
	// Create handlers
	carHandler := handlers.NewCarHandler()

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "message": "Race Cars API is running"}`))
	}).Methods("GET")

	// Root endpoint
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"name": "Race Cars API",
			"version": "1.0.0",
			"description": "A RESTful API for managing race cars",
			"endpoints": {
				"cars": "/api/cars",
				"health": "/health"
			}
		}`))
	}).Methods("GET")
} 