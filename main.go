package main

import (
	"time"

	"github.com/mgmaster24/m2-pokedex/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	startREPL(&config{
		client:        pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	})
}
