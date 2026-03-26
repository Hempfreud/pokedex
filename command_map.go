package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config, arg string) error {
	Locations, err := c.pokeapiClient.ListLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}
	c.nextLocationsURL = Locations.Next
	c.prevLocationsURL = Locations.Previous

	for _, location := range Locations.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func commandMapb(c *config, arg string) error {
	if c.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	Locations, err := c.pokeapiClient.ListLocations(c.prevLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = Locations.Next
	c.prevLocationsURL = Locations.Previous

	for _, location := range Locations.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}
