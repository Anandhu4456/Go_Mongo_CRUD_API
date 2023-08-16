package routes

import (
	"go-mongo/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	route.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	route.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	route.HandleFunc("/api/movie/{id}", controller.DltOneMovie).Methods("DELETE")
	route.HandleFunc("/api/deleteallmovie", controller.DltAllMovie).Methods("DELETE")

	return route
}
