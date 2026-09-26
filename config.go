package main

type config struct {
	commands map[string]pokedexCommand
	Next     string `json:"next"`
	Previous string `json:"previous"`
}
