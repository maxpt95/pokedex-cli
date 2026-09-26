package main

import "github.com/maxpt95/pokedexcli/internal/pokeapi"

type config struct {
	commands       map[string]pokedexCommand
	client         pokeapi.Client
	pokeApiNextUrl *string
	pokeApiPrevUrl *string
}
