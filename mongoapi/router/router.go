package router

import (
	"fmt"

	moviecontroller "github.com/Yagnik007/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	fmt.Println("Router")

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", moviecontroller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", moviecontroller.CreateOneMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", moviecontroller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", moviecontroller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", moviecontroller.DeleteAllMovies).Methods("DELETE")

	return router
}
