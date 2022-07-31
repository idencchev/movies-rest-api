package router

import (
	"movies-rest-api/services"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", services.GetAllMovies).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/movies/search", services.SearchMovies).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/users", services.RegisterUser).Methods("POST", "OPTIONS")

	return router
}
