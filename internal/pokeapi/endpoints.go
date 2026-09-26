package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) reqShallowList(requestUrl *string) (ShallowList, error) {
	resp, err := c.httpClient.Get(*requestUrl)
	if err != nil {
		return ShallowList{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ShallowList{}, err
	}

	var shallowList ShallowList
	if err := json.Unmarshal(data, &shallowList); err != nil {
		return ShallowList{}, err
	}

	return shallowList, nil
}

func (c *Client) ListLocationAreas(requestUrl *string) (ShallowList, error) {
	url := pokeApiUrl + "location-area/"
	if requestUrl != nil {
		url = *requestUrl
	}

	locationAreas, err := c.reqShallowList(&url)
	if err != nil {
		return ShallowList{}, fmt.Errorf("error: failed to retrieve locaction areas %w", err)
	}

	return locationAreas, nil
}
