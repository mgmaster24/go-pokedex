package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) GetLocations(url string) (*PokeResponse, error) {
	if val, ok := client.cache.Get(url); ok {
		pokeResponse := PokeResponse{}
		err := json.Unmarshal(val, &pokeResponse)
		if err != nil {
			return nil, err
		}

		return &pokeResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating locations request for PokeAPI, Error: %v", err)
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting locations from PokeAPI, Error: %v", err)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response body, Error: %v", err)
	}

	pokeResponse := PokeResponse{}
	err = json.Unmarshal(data, &pokeResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling the response data, Error: %v", err)
	}

	client.cache.Add(url, data)
	return &pokeResponse, nil
}
