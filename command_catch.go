package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, arg string) error {
	if arg == "" {
		return fmt.Errorf("please specify a Pokemon to catch")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	//Get Pokemon data from pokeapi
	pokemon, err := c.pokeapiClient.GetPokemon(arg)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)
	if res > 50 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		c.caughtPokemon[pokemon.Name] = pokemon
	}
	return nil
}
