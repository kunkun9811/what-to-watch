package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func createRouter() http.Handler {
	getAllAnime := GetAllAnimeHandler{}
	searchAnimeByName := SearchAnimeByName{}

	r := mux.NewRouter()

	// TODO: about to add handler
	r.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
	})
	animeRouter := r.PathPrefix("/api/anime").Subrouter()
	animeRouter.Handle("", getAllAnime).Methods("GET")
	animeRouter.Handle("", searchAnimeByName).Methods("POST")
	return r
}
