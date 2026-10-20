package main

import "fmt"

func commandPokedex(conf *config, args ...string) error {
	if len(conf.Pokedex) == 0 {
		return fmt.Errorf("You have not caught any Pokemon!")
	}
	fmt.Println("Your Pokedex:")
	for key := range conf.Pokedex {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
