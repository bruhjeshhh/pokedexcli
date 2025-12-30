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
		
		}
	}
}
