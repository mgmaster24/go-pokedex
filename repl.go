package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL(cfg *config) {
	fmt.Println("------ M2 Pokedex -------")
	fmt.Println("-------------------------")
	cmdScanner := bufio.NewScanner(os.Stdin)
	cmds := getCommands()
	for {
		fmt.Print("Pokedex ->")
		cmdScanner.Scan()

		words := cleanInput(cmdScanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, found := cmds[commandName]
		if found {
			err := cmd.callback(cfg, args...)
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
