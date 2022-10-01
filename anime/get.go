package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type GetAllAnimeHandler struct{}

type GetAllAnimeResponse struct {
	AnimeList []string `json:"animeList"`
}

func (g GetAllAnimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	response := GetAllAnimeResponse{
		AnimeList: animeWannaWatch,
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(os.Stdout).Encode(&response)
	_ = json.NewEncoder(w).Encode(&response)
}
