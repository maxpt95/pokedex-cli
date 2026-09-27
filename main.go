package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/maxpt95/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{commands: getCommands(), client: pokeapi.NewClient(5*time.Second, 10*time.Second)}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		fmt.Println(cleanedInput)

		command, ok := cfg.commands[cleanedInput[0]]
		if !ok {
			fmt.Println("Uknown command")
			continue
		}
		params := cleanedInput[1:]
		command.callback(&cfg, params)
	}

}
