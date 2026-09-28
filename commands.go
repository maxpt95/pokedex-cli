package main

import (
	"fmt"
	"os"
)

type pokedexCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"catch": {
			name:        "catch",
			description: "Throw a pokeball and try to catch a Pokemon!",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range cfg.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config, args ...string) error {
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

func commandMapb(cfg *config, args ...string) error {
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

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expecting 1 parameter <locationName> and got %d", len(args))
	}

	locationName := args[0]

	locationArea, err := cfg.client.LocationArea(locationName)
	if locationArea.ID == 0 {
		return fmt.Errorf("location %s can't be found", locationArea)
	}
	fmt.Printf("Exploring %s \n", locationName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expecting 1 parameter <pokemonName> and got %d", len(args))
	}
	pokemonName := args[0]

	pokemon, err := cfg.client.Pokemon(pokemonName)
	if err != nil {
		return err
	}
	if pokemon.ID == 0 {
		return fmt.Errorf("pokemon %s doesn't exist", pokemonName)
	}

	fmt.Println("Throwing a Pokeball to %s...", pokemonName)

	return nil
}
