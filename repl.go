package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var naksha = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Provide help to the user",
		callback:    sendHelp,
	},
}

func cleanInput(text string) []string {
	// var opt []string

	words := strings.Fields(text)

	return words

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("exit")

}

func sendHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return errors.New("exit")
}
