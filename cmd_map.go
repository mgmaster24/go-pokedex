package main

import (
	"github.com/mgmaster24/m2-pokedex/pokeapi"
)

func map_cmd(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != "" {
		url = cfg.Next
	}

	resp, err := pokeapi.GetLocations(url)
	if err != nil {
		return err
	}

	resp.Print()
	cfg.Next = resp.Next
	cfg.Prev = resp.Previous
	return nil
}
