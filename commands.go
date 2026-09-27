package main

import (
	"errors"
	"fmt"
	"os"
)

type pokedexCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
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
		"explore": {
			name:        "explore",
			description: "Explore a location for Pokemon",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
func commandExit(cfg *config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, params []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range cfg.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config, params []string) error {
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

func commandMapb(cfg *config, params []string) error {
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

func commandExplore(cfg *config, params []string) error {
	if len(params) != 1 {
		errorMsg := fmt.Sprintf("error: expecting 1 parameter and got %d", len(params))
		fmt.Println(errorMsg)
		return errors.New(errorMsg)
	}

	locationArea, err := cfg.client.LocationArea(params[0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
