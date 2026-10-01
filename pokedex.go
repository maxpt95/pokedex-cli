package main

import (
	"fmt"

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
