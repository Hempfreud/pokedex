package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hempfreud/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

var commands map[string]cliCommand

func cleanInput(text string) []string {
	// split the user's input into "words" based on whitespace.
	// It should also lowercase the input and trim any leading or trailing whitespace
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}

func startRepl(c *config) {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next location",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous location",
			callback:    commandMapb,
		},
	}
	// Create a bufio.Scanner that reads from os.Stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Create infinite Loop
	for {
		// Print the prompt
		fmt.Print("Pokedex > ")

		// Read the user input
		scanner.Scan()
		input := scanner.Text()

		// Clean the user's input string
		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}
		// Update your REPL loop to use the "command" the user typed in to look up the callback function in the registry.
		// If the command is found, call the callback (and print any errors that are returned).
		if cmd, ok := commands[cleanedInput[0]]; ok {
			err := cmd.callback(c)
			if err != nil {
				fmt.Printf("Error executing command: %s\n", err)
			}
		} else {
			fmt.Printf("Unknown command")
		}
	}
}
