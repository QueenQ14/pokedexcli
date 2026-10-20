package main

import "fmt"

func commandExplore(conf *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide a location area name!")
	}
	areaName := args[0]
	res, err := conf.pokeapiClient.ListPokemonArea(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s\n", areaName)
	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
