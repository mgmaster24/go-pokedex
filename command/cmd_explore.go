package command

import (
	"errors"
	"fmt"

	"github.com/mgmaster24/m2-pokedex/config"
)

func exploreLocation(config *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + name
	location, err := config.Client.ExploreLocation(url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
