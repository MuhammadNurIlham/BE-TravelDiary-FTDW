package routes

import (
	"be-journey/handlers"
	"be-journey/pkg/middleware"
	"be-journey/pkg/mysql"
	"be-journey/repositories"

	"github.com/gorilla/mux"
)

func BookmarkRoutes(r *mux.Router) {
	bookmarkRepository := repositories.RepositoryBookmark(mysql.DB)
	h := handlers.HandlerBookmark(bookmarkRepository)

	r.HandleFunc("/bookmarks", h.FindBookmarks).Methods("GET") //getAll
	r.HandleFunc("/bookmark/{id}", h.GetBookmark).Methods("GET") //select
	r.HandleFunc("/bookmark", middleware.Auth(h.CreateBookmark)).Methods("POST") //all
	r.HandleFunc("/bookmark/{id}", middleware.Auth(h.DeleteBookmark)).Methods("DELETE") //delete select
	r.HandleFunc("/bookmarks/{id}", middleware.Auth(h.DeleteBookmarks)).Methods("DELETE") //delete select
	r.HandleFunc("/bookmarks/{id}", h.GetBookmarks).Methods("GET") //select
}
