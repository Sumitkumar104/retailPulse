package routes

import (
	"github.com/gorilla/mux"
	"retailPulse/handlers"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/submit/", handlers.SubmitJob).Methods("POST")
	router.HandleFunc("/api/status", handlers.GetJobStatus).Methods("GET")

	return router
}
