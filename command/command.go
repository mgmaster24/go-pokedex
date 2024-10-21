package command

import (
	"errors"

	"github.com/mgmaster24/m2-pokedex/config"
)

type Command struct {
	name        string
	description string
	Callback    func(cfg *config.Config, args ...string) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    help,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			Callback:    exit,
		},
		"map": {
			name:        "map",
			description: "returns the next 20 name of locations in the Pokemon world",
			Callback:    map_cmd,
		},
		"mapb": {
			name:        "mapb",
			description: "returns the previous 20 names of locations in the Pokemon world",
			Callback:    mapb_cmd,
		},
		"explore": {
			name:        "explore",
			description: "returns a list of pokemon located in the provided location",
			Callback:    exploreLocation,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch a pokemon by name",
			Callback:    catch,
		},
		"inspect": {
			name:        "inspect",
			description: "display details about caught pokemon",
			Callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list the caught pokemon",
			Callback:    pokedex,
		},
	}
}

func GetArgs(numArgs int, args ...string) ([]string, error) {
	if len(args) != numArgs {
		return nil, errors.New("Not enough arguments for this command")
	}

	return args[:numArgs], nil
}
