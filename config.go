package main

import "github.com/mgmaster24/m2-pokedex/pokeapi"

type config struct {
	Next          *string
	Prev          *string
	client        pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
}
