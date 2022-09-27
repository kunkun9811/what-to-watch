package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func createRouter() http.Handler {
	r := mux.NewRouter()

	// TODO: about to add handler
	animeRouter := r.PathPrefix("/api/anime").Subrouter()
	animeRouter.Handle("").Methods("GET")
}
