package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

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
