package main

import (
	"fmt"
)

func commandMapF(conf *config, args ...string) error {
	res, err := conf.pokeapiClient.GetLocationArea(conf.Next)
	if err != nil {
		return err
	}
	conf.Next = res.Next
	conf.Previous = res.Previous

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapB(conf *config, args ...string) error {
	if conf.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}
	res, err := conf.pokeapiClient.GetLocationArea(conf.Previous)
	if err != nil {
		return err
	}
	conf.Next = res.Next
	conf.Previous = res.Previous

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}
	return nil
}
