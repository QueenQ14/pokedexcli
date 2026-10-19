package main

import "fmt"

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	for key, val := range getCommands() {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	return nil
}
