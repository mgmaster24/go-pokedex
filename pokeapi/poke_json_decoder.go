package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func decodePokeJSON(resp *http.Response) (*PokeResponse, error) {
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response body, Error: %v", err)
	}

	pokeResp := PokeResponse{}
	err = json.Unmarshal(data, &pokeResp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling the response data, Error: %v", err)
	}

	return &pokeResp, nil
}
