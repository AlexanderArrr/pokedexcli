package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AlexanderArrr/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	Next     int
	Previous int
}

func REPL() {
	commands := createCommands()
	cfg := config{
		Next:     1,
		Previous: 1,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		output := cleanInput(input)
		if _, ok := commands[output[0]]; !ok {
			fmt.Println("Unknown command")
		} else {
			commands[output[0]].callback(&cfg)
		}
	}
}

func createCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show the next 20 available location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 available location areas",
			callback:    commandMapb,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	commands := createCommands()

	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	location_area := pokeapi.LocationArea{}
	for i := cfg.Next; i < (cfg.Next + 20); i++ {
		url := "https://pokeapi.co/api/v2/location-area/" + strconv.Itoa(i)
		err := pokeapi.MakePokeAPIRequest(url, &location_area)
		if err != nil {
			fmt.Println("Error within commandMap(): %w", err)
		}
		fmt.Println(location_area.Name)
	}
	cfg.Next += 20
	cfg.Previous = cfg.Next - 20

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == 1 {
		fmt.Println("you're on the first page")
		return nil
	}
	location_area := pokeapi.LocationArea{}
	for i := (cfg.Previous - 20); i < cfg.Previous; i++ {
		url := "https://pokeapi.co/api/v2/location-area/" + strconv.Itoa(i)
		err := pokeapi.MakePokeAPIRequest(url, &location_area)
		if err != nil {
			fmt.Println("Error within commandMapb(): %w", err)
		}
		fmt.Println(location_area.Name)
	}
	cfg.Next -= 20
	cfg.Previous = cfg.Next - 20

	return nil
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
