package command

import "github.com/mgmaster24/m2-pokedex/config"

func map_cmd(cfg *config.Config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	resp, err := cfg.Client.GetLocations(url)
	if err != nil {
		return err
	}

	resp.Print()
	cfg.Next = resp.Next
	cfg.Prev = resp.Previous
	return nil
}
