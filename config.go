package main

import (
	"github.com/QueenQ14/pokedexcli/internal/pokeapi"
)

type config struct {
	supportedCommands map[string]cliCommand
	pokeapiClient     pokeapi.Client
	Next              *string
	Previous          *string
	Pokedex           map[string]pokeapi.ResponsePokemonInfo
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
		"explore": {
			name:        "explore <location-name>",
			description: "List all Pokemon encounters in given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Attempt to catch the given Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Inspect the details of your caught Pokemon",
			callback:    commandInspect,
		},
	}
}
