package command

import (
	"fmt"
	"os"

	"github.com/mgmaster24/m2-pokedex/config"
)

func exit(cfg *config.Config, args ...string) error {
	fmt.Println("Exiting Pokedex")
	os.Exit(0)
	return nil
}
