package command

import (
	"errors"
	"fmt"

	"github.com/mgmaster24/m2-pokedex/config"
)

func pokedex(config *config.Config, args ...string) error {
	if len(args) > 0 {
		return errors.New("this command takes no arguments")
	}

	fmt.Println("Your Pokedex:")
	for _, p := range config.CaughtPokemon {
		fmt.Printf("\t- %s\n", p.Name)
	}

	return nil
}
