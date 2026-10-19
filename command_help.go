package main

import "fmt"

func commandHelp(conf *config) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	for key, val := range conf.supportedCommands {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	return nil
}
