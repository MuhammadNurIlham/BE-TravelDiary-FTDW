package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	JourneyRoutes(r)
	BookmarkRoutes(r)
	AuthRoutes(r)
}
