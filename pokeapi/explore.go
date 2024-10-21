package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) ExploreLocation(locationUrl string) (Location, error) {
	if val, ok := client.cache.Get(locationUrl); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", locationUrl, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, err
	}

	client.cache.Add(locationUrl, data)
	return locationResp, nil
}
