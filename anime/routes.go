package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func createRouter() http.Handler {
	getAllAnime := GetAllAnimeHandler{}

	r := mux.NewRouter()

	// TODO: about to add handler
	r.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
	})
	animeRouter := r.PathPrefix("/api/anime").Subrouter()
	animeRouter.Handle("", getAllAnime).Methods("GET")
	return r
}
