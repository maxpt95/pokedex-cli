// A pokeapi client implemented following the documentation of pokeapi endpoints
// https://pokeapi.co/docs/v2#resource-listspagination-section
package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const pokeApiUrl = "https://pokeapi.co/api/v2/"

type Named struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getNamed(requestUrl string) (Named, error) {
	resp, err := http.Get(requestUrl)
	if err != nil {
		return Named{}, err
	}
	var named Named
	if err := json.NewDecoder(resp.Body).Decode(&named); err != nil {
		return Named{}, err
	}

	return named, nil
}

func GetLocationAreas(requestUrl string) (Named, error) {
	if requestUrl == "" {
		requestUrl = pokeApiUrl + "location-area/"
	}

	locationAreas, err := getNamed(requestUrl)
	if err != nil {
		return Named{}, fmt.Errorf("error: failed to retrieve locaction areas %w", err)
	}

	return locationAreas, nil
}
