package main

import "fmt"

func help(cfg *config) error {
	fmt.Println("Welcome to the M2 Pokedex")
	fmt.Println("Usage:")

	cmds := getCommands()
	for k, v := range cmds {
		fmt.Printf("%s: %s\n", k, v.description)
	}

	return nil
}
