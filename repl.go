package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replStart(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		if len(cleanedInput) < 1 {
			continue
		}

		command, ok := cfg.commands[cleanedInput[0]]
		if !ok {
			fmt.Println("Uknown command")
			continue
		}

		args := cleanedInput[1:]
		command.callback(cfg, args...)
	}
}
func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	cleaned := strings.Fields(lowered)

	return cleaned

}
