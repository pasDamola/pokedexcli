package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() string
}

func commandHelp() string {
 return  `
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`
}

func commandExit() string {
	return "exit"
}





func main() {

	 command := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	for {

		// Create a new scanner to read from standard input (os.Stdin)
		scanner := bufio.NewScanner(os.Stdin)
   
		fmt.Print("Pokedex > ")
	
		// Scan for input
		if scanner.Scan() {
			// Use .Text() to get the input as a string
			inputText := scanner.Text()

			_, ok := command[inputText]

			if ok {
				s := command[inputText].callback()
				
				if s == "exit" {
					break
				}

				fmt.Println(s)

			} else {
				fmt.Printf("%s is not a valid command in the program\n", inputText)
			}
			
		} else {
			fmt.Println("An error occurred while reading input.")
		}
	}
}