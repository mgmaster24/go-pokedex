package main

import (
	"fmt"

	"github.com/mgmaster24/m2-pokedex/pokeapi"
)

func mapb_cmd(cfg *config) error {
	if cfg.Prev != "" {
		resp, err := pokeapi.GetLocations(cfg.Prev)
		if err != nil {
			return err
		}

		resp.Print()
		return nil
	}

	fmt.Println("no previous result to retrieve")
	return nil
}
