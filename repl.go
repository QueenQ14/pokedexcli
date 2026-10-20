package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			scanner.Err()
		}
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		cmdName := words[0]
		args := words[1:]
		if command, ok := conf.supportedCommands[cmdName]; ok {
			err := command.callback(conf, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
