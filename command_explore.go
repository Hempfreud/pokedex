package main

import (
	"fmt"
)

func commandExplore(c *config, arg string) error {
	fmt.Printf("Exploring %s...\n", arg)

	result, err := c.pokeapiClient.GetLocationAreaDetails(arg)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	// loop over result.PokemonEncounters and print each name

	for _, encounter := range result.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
