package main

import (
	"encoding/json"
	"net/http"
)

type GetAllMoviesHandler struct{}

type GetAllMoviesResponse struct {
	MoviesList []string `json:"moviesList"`
}

func (g GetAllMoviesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := GetAllMoviesResponse{
		MoviesList: moviesWannaWatch,
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
