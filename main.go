package main

import (
	"time"

	"github.com/maxpt95/pokedexcli/internal/pokeapi"
)

func main() {
	pokedex := Pokedex{caughtPokemons: make(map[string]pokeapi.Pokemon)}
	cfg := config{commands: getCommands(), client: pokeapi.NewClient(5*time.Second, 10*time.Second), pokedex: pokedex}
	startRepl(&cfg)

}
