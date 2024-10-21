package main

import (
	"fmt"
	"os"
)

func exit(cfg *config, args ...string) error {
	fmt.Println("Exiting Pokedex")
	os.Exit(0)
	return nil
}
