package main

import (
	"fmt"
)

func mapb_cmd(cfg *config, args ...string) error {
	if cfg.Prev != nil {
		resp, err := cfg.client.GetLocations(*cfg.Prev)
		if err != nil {
			return err
		}

		resp.Print()
		return nil
	}

	fmt.Println("no previous result to retrieve")
	return nil
}
