package main

import (
	"fmt"
	"os"
)

func exit(cfg *config) error {
	fmt.Println("Exiting Pokedex")
	os.Exit(0)
	return nil
}
