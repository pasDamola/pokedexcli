package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	for {

		// Create a new scanner to read from standard input (os.Stdin)
		scanner := bufio.NewScanner(os.Stdin)
   
		fmt.Print("Pokedex > ")
	
		// Scan for input
		if scanner.Scan() {
			// Use .Text() to get the input as a string
			inputText := scanner.Text()
			fmt.Printf("You entered: %s\n", inputText)
		} else {
			fmt.Println("An error occurred while reading input.")
		}
	}
}