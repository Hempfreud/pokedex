package main

import (
	"fmt"
)

func commandPokedex(c *config, arg string) error {
	if len(c.caughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range c.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
