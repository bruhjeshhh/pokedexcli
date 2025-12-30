package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex >")
		reader.Scan()

		input := reader.Text()
		value, ok := naksha[input]
		if ok == false {
			fmt.Println("Unknown command")
		} else {
			_ = value.callback()

		}
	}
}
