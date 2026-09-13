package main

type pokedexCommand struct {
	name        string
	description string
	callback    func() error
}

var pokedexCommands = map[string]pokedexCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}
