package main

import "github.com/maxpt95/pokedexcli/internal/pokeapi"

type Pokedex struct {
	caughtPokemons map[string]pokeapi.Pokemon
}

func (dex Pokedex) Add(pokemon pokeapi.Pokemon) {
	dex.caughtPokemons[pokemon.Name] = pokemon
}
