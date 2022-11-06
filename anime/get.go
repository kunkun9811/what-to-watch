package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GetAllAnimeHandler struct{}

type GetAllAnimeResponse struct {
	AnimeList []string `json:"animeList"`
}

func (g GetAllAnimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := GetAllAnimeResponse{
		AnimeList: animeWannaWatch,
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(os.Stdout).Encode(&response)
	_ = json.NewEncoder(w).Encode(&response)
}

type SearchAnimeByNameRequest struct {
	Name string `json:"name" validate:"required,oneof=DEMON_SLAYER SPY_X_FAMILY TOKYO_REVENGERS JUJUTSU_KAISEN_0 THE_PRINCESS ZERO_TWO"`
}

type SearchAnimeByNameResponse struct {
	Anime string `json:"anime"`
}

type SearchAnimeByName struct{}

func (g SearchAnimeByName) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var searchAnimeByNameRequest = new(SearchAnimeByNameRequest)
	if err := ValidateAndPopulateRequestBody(w, r, searchAnimeByNameRequest); err != nil {
		return
	}

	anime, exist := animeWannaWatchMap[searchAnimeByNameRequest.Name]
	if !exist {
		errorTxt := fmt.Sprintf("Anime with name '%v' does not exist", searchAnimeByNameRequest.Name)
		fmt.Printf("Error: %v\n", errorTxt)
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ValidationError{
			Errors: map[string][]string{
				"Name": {errorTxt},
			},
		})
		return
	}

	response := SearchAnimeByNameResponse{
		Anime: anime,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(os.Stdout).Encode(&response)
	_ = json.NewEncoder(w).Encode(&response)
}
