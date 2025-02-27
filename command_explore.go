package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	locationArea, err := cfg.pokeapiClient.LocationArea(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationArea.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
