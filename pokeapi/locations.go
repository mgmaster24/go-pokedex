package pokeapi

import (
	"fmt"
	"net/http"
)

func GetLocations(url string) (*PokeResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting locations from PokeAPI, Error: %v", err)
	}

	pokeResponse, err := decodePokeJSON(resp)
	if err != nil {
		return nil, fmt.Errorf("error decoding the poke response for locations, Error: %v", err)
	}

	return pokeResponse, nil
}
