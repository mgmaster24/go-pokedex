package main

func map_cmd(cfg *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	resp, err := cfg.client.GetLocations(url)
	if err != nil {
		return err
	}

	resp.Print()
	cfg.Next = resp.Next
	cfg.Prev = resp.Previous
	return nil
}
