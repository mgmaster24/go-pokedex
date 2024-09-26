package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	fmt.Println("------ M2 Pokedex -------")
	fmt.Println("-------------------------")
	cmdScanner := bufio.NewScanner(os.Stdin)
	cmds := getCommands()
	cfg := config{}
	for {
		fmt.Print("Pokedex ->")
		cmdScanner.Scan()

		words := cleanInput(cmdScanner.Text())
		if cmd, found := cmds[words[0]]; found {
			err := cmd.callback(&cfg)
			if err != nil {
				fmt.Printf("error running cmd, err: %v\n", err)
			}
		} else {
			fmt.Println("command not found")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
