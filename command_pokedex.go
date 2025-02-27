package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("you have not caught any pokemon yet")
	}

	fmt.Println("Your Pokedex:")
	for p, _ := range cfg.pokedex {
		fmt.Printf("  - %s\n", p)
	}

	return nil
}
