package main

import (
	"errors"
	"fmt"
	"maps"
	"slices"

	"github.com/maxpt95/pokedexcli/internal/pokeapi"
)

type Pokedex struct {
	caughtPokemons map[string]pokeapi.Pokemon
}

func (dex Pokedex) Add(pokemon pokeapi.Pokemon) {
	dex.caughtPokemons[pokemon.Name] = pokemon
}

func (dex Pokedex) Get(pokemonName string) (pokeapi.Pokemon, error) {
	pokemon, ok := dex.caughtPokemons[pokemonName]
	if !ok {
		return pokeapi.Pokemon{}, fmt.Errorf("you have not caught %s", pokemonName)
	}
	return pokemon, nil
}

func (dex Pokedex) List() ([]string, error) {
	if len(dex.caughtPokemons) < 1 {
		return nil, errors.New("your Pokedex is empty!")
	}
	return slices.Sorted(maps.Keys(dex.caughtPokemons)), nil
}
