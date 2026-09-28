package main

import (
	"time"

	"github.com/maxpt95/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{commands: getCommands(), client: pokeapi.NewClient(5*time.Second, 10*time.Second)}
	startRepl(&cfg)

}
