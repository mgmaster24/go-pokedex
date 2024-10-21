package config

import "github.com/mgmaster24/m2-pokedex/pokeapi"

type Config struct {
	Next          *string
	Prev          *string
	Client        pokeapi.Client
	CaughtPokemon map[string]pokeapi.Pokemon
}
