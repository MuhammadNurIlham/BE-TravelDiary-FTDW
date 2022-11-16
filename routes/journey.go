package routes

import (
	"be-journey/handlers"
	"be-journey/pkg/middleware"
	"be-journey/pkg/mysql"
	"be-journey/repositories"

	"github.com/gorilla/mux"
)

func JourneyRoutes(r *mux.Router) {
	journeyRepository := repositories.RepositoryJourney(mysql.DB)
	h := handlers.HandlerJourney(journeyRepository)

	r.HandleFunc("/journeys", h.FindJourneys).Methods("GET")
	// r.HandleFunc("/journeys/{user_id}", h.GetJourney).Methods("GET") // for get journey base on user id
	r.HandleFunc("/journey/{id}", h.GetJourney).Methods("GET") // for get journey base on id journey
	r.HandleFunc("/journey", middleware.Auth(middleware.UploadFile(h.CreateJourney))).Methods("POST")
	// r.HandleFunc("/journeys/{user_id}", h.DeleteJourney).Methods("DELETE")
	r.HandleFunc("/journey/{id}", middleware.Auth(h.UpdateJourney)).Methods("PATCH") //for update
	r.HandleFunc("/journey/{id}", h.DeleteJourney).Methods("DELETE")
}
