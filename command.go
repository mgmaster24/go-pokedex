package main

type command struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    help,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    exit,
		},
		"map": {
			name:        "map",
			description: "returns the next 20 name of locations in the Pokemon world",
			callback:    map_cmd,
		},
		"mapb": {
			name:        "mapb",
			description: "returns the previous 20 names of locations in the Pokemon world",
			callback:    mapb_cmd,
		},
		"explore": {
			name:        "explore",
			description: "returns a list of pokemon located in the provided location",
			callback:    exploreLocation,
		},
	}
}
