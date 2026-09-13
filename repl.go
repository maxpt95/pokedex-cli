package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	cleaned := strings.Fields(lowered)

	return cleaned

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
