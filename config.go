package main

import "github.com/QueenQ14/pokedexcli/internal/pokeapi"

type config struct {
	supportedCommands map[string]cliCommand
	pokeapiClient     pokeapi.Client
	Next              *string
	Previous          *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas; then moves the next index",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays 20 location areas; then moves the previous index",
			callback:    commandMapB,
		},
	}
}
