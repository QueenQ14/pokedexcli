package main

import (
	"time"

	"github.com/QueenQ14/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 60*time.Second)
	conf := &config{
		supportedCommands: getCommands(),
		pokeapiClient:     pokeClient,
		Pokedex:           make(map[string]pokeapi.ResponsePokemonInfo),
	}
	startRepl(conf)
}
