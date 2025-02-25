package main

import (
	"time"

	"github.com/AlexanderArrr/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	REPL(cfg)
}
