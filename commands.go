package main

import (
	"fmt"
	"os"
)

type pokedexCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]pokedexCommand {
	return map[string]pokedexCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next location areas of the pokemon worlds",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous location areas of the pokemon worlds",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range cfg.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	locationAreas, err := cfg.client.ListLocationAreas(cfg.pokeApiNextUrl)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	cfg.pokeApiNextUrl = locationAreas.Next
	cfg.pokeApiPrevUrl = locationAreas.Previous

	return nil

}

func commandMapb(cfg *config) error {
	if cfg.pokeApiPrevUrl == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locationAreas, err := cfg.client.ListLocationAreas(cfg.pokeApiPrevUrl)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	cfg.pokeApiNextUrl = locationAreas.Next
	cfg.pokeApiPrevUrl = locationAreas.Previous

	return nil

}
