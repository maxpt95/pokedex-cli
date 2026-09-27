package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) makeRequest(requestUrl *string) ([]byte, error) {
	if data, ok := c.cache.Get(*requestUrl); ok {
		return data, nil
	}

	resp, err := c.httpClient.Get(*requestUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(*requestUrl, data)
	return data, nil
}
func (c *Client) reqShallowList(requestUrl *string) (ShallowList, error) {
	data, err := c.makeRequest(requestUrl)
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

func (c *Client) LocationArea(locationName string) (LocationArea, error) {
	url := pokeApiUrl + "location-area/" + locationName
	data, err := c.makeRequest(&url)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error: failed to retrieve location area <%s> %w", locationName, err)
	}

	var locationArea LocationArea
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationArea{}, fmt.Errorf("error: failed to parse location area json %w", locationName, err)
	}

	return locationArea, nil
}
