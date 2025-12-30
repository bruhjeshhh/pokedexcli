package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	url := "https://pokeapi.co/api/v2/location-area"
	urlHolder := config{
		nextUrl: &url,
	}
	for true {
		fmt.Print("Pokedex >")
		reader.Scan()

		input := reader.Text()
		value, ok := naksha[input]
		if ok == false {
			fmt.Println("Unknown command")
		} else {
			_ = value.callback(&urlHolder)

		}
	}
}
