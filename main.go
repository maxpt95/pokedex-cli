package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cfg := config{commands: getCommands()}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		command, ok := cfg.commands[cleanedInput[0]]
		if !ok {
			fmt.Println("Uknown command")
			continue
		}
		command.callback(&cfg)
	}

}
