// A pokeapi client implemented following the documentation of pokeapi endpoints
// https://pokeapi.co/docs/v2#resource-listspagination-section
package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const pokeApiUrl = "https://pokeapi.co/api/v2/"

func getNamed(requestUrl string) (ShallowList, error) {
	resp, err := http.Get(requestUrl)
	if err != nil {
		return ShallowList{}, err
	}
	var named ShallowList
	if err := json.NewDecoder(resp.Body).Decode(&named); err != nil {
		return ShallowList{}, err
	}

	return named, nil
}

func GetLocationAreas(requestUrl string) (ShallowList, error) {
	if requestUrl == "" {
		requestUrl = pokeApiUrl + "location-area/"
	}

	locationAreas, err := getNamed(requestUrl)
	if err != nil {
		return ShallowList{}, fmt.Errorf("error: failed to retrieve locaction areas %w", err)
	}

	return locationAreas, nil
}
