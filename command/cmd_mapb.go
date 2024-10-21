package command

import (
	"fmt"

	"github.com/mgmaster24/m2-pokedex/config"
)

func mapb_cmd(cfg *config.Config, args ...string) error {
	if cfg.Prev != nil {
		resp, err := cfg.Client.GetLocations(*cfg.Prev)
		if err != nil {
			return err
		}

		resp.Print()
		return nil
	}

	fmt.Println("no previous result to retrieve")
	return nil
}
