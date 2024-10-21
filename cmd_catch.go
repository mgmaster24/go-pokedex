package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	pokemon, err := config.client.Catch(url)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)
	pokemonName := pokemon.Name
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if res > 42 {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	config.caughtPokemon[pokemonName] = pokemon
	return nil
}
