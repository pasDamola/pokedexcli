package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

type PokeAPI struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandHelp() {
 fmt.Println(  `
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
}

func commandExit()  {

}

func getNextLocations()  {
	
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location?offset=%v&limit=20", )
	resp, err := http.Get(endpoint)
	if err != nil {
		// handle error
		log.Fatalf("couldn't read request because %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Fatalf("couldn't read request because %v", err)
	}

	pokeAPI := PokeAPI{}

	if json.Unmarshal(body, &pokeAPI); err != nil {
		log.Fatalf("couldn't read request because %v", err)
	}

	next := pokeAPI.Next
	results := pokeAPI.Results

	for _, val := range results {
		fmt.Println(val.Name)
	}
}


func getPrevLocations()  {
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
		"map": {
			name: "map",
			description: "Get next 20 locations from the PokeAPI",
			callback: getNextLocations,
		},
		"mapb": {
			name: "mapb",
			description: "Get previous 20 locations from the PokeAPI",
			callback: getPrevLocations,
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
				command[inputText].callback()
				

			} else {
				fmt.Printf("%s is not a valid command in the program\n", inputText)
			}
			
		} else {
			fmt.Println("An error occurred while reading input.")
		}
	}
}