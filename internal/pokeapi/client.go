package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const BaseURL = "https://pokeapi.co/api/v2/location-area/"

type Location struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

func FetchLocation(url string) (*Location, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if req.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d\n", req.StatusCode)
	}

	data, err := io.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	var newLocation Location
	if err := json.Unmarshal(data, &newLocation); err != nil {
		return nil, err
	}
	return &newLocation, nil
}

func FetchLocationArea(url string) (*LocationAreaResponse, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if req.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d\n", req.StatusCode)
	}

	data, err := io.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	var newLocation LocationAreaResponse
	if err := json.Unmarshal(data, &newLocation); err != nil {
		return nil, err
	}
	return &newLocation, nil
}
