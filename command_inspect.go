package main

import "fmt"

func commandInspect(conf *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide a Pokemon name!")
	}
	pkmnName := args[0]
	val, ok := conf.Pokedex[pkmnName]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	stats := ""
	for _, stat := range val.Stats {
		entry := fmt.Sprintf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		stats += entry
	}
	types := ""
	for _, typing := range val.Types {
		entry := fmt.Sprintf("  - %s\n", typing.Type.Name)
		types += entry
	}
	fmt.Printf(`Name: %s
Height: %d
Weight: %d
Stats:
%sTypes:
%s`, val.Name, val.Height, val.Weight, stats, types)

	return nil

}
