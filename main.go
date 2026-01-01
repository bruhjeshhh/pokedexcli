package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	url := "https://pokeapi.co/api/v2/location-area/"
	urlHolder := config{
		nextUrl: &url,
	}
	for true {
		fmt.Print("Pokedex >")
		reader.Scan()

		input := reader.Text()
		words := strings.Fields(input)
		value, ok := naksha[words[0]]
		if ok == false {
			fmt.Println("Unknown command")
		} else {
			if len(words) == 2 {
				_ = value.callback(&urlHolder, words[1])
			} else {
				_ = value.callback(&urlHolder, "")
			}

		}
	}
}
