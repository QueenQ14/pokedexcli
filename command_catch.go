package main

import (
	"fmt"
	"math/rand/v2"
)

const (
	min_exp = 30.00
	max_exp = 600.00
)

func commandCatch(conf *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide a Pokemon name!")
	}
	pkmnName := args[0]
	res, err := conf.pokeapiClient.GetPokemonInfo(pkmnName)
	if err != nil {
		return err
	}
	C := max_exp * (1 - ((float64(res.BaseExperience) - min_exp) / (max_exp - min_exp)))
	catchRate := max(min_exp, min(max_exp, C))
	roll := rand.IntN(max_exp) + 1
	fmt.Printf("Throwing a Pokeball at %s...\n", res.Name)
	if float64(roll) <= catchRate {
		conf.Pokedex[res.Name] = res
		fmt.Printf("%s was caught!\n", res.Name)
		return nil
	} else {
		fmt.Printf("%s escaped!\n", res.Name)
		return nil
	}
}
