package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func createRouter() http.Handler {
	getAllMoviesHandler := GetAllMoviesHandler{}

	r := mux.NewRouter()

	moviesRouter := r.PathPrefix("/api/movies").Subrouter()
	moviesRouter.Handle("", getAllMoviesHandler).Methods("GET")
	return r
}
