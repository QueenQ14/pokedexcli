package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			scanner.Err()
		}
		words := cleanInput(scanner.Text())
		cmdName := words[0]
		fmt.Printf("Your command was: %s\n", cmdName)
	}
}
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
