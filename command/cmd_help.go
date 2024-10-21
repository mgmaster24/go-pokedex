package command

import (
	"fmt"
	"sort"

	"github.com/mgmaster24/m2-pokedex/config"
)

func help(cfg *config.Config, args ...string) error {
	fmt.Println("Welcome to the M2 Pokedex")
	fmt.Println("Usage:")

	cmds := GetCommands()
	// Extract keys
	keys := make([]string, 0, len(cmds))
	for k := range cmds {
		keys = append(keys, k)
	}

	// Sort keys
	sort.Strings(keys)

	// Iterate through sorted keys
	for _, k := range keys {
		fmt.Printf("%s: %s\n", k, cmds[k].description)
	}

	return nil
}
