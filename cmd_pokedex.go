package main

import (
	"errors"
	"fmt"
)

func pokedex(config *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("this command takes no arguments")
	}

	fmt.Println("Your Pokedex:")
	for _, p := range config.caughtPokemon {
		fmt.Printf("\t- %s\n", p.Name)
	}

	return nil
}
