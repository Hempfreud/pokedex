package main

import (
	"time"

	"github.com/hempfreud/pokedex/internal/pokeapi"
	"github.com/hempfreud/pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	client := pokeapi.NewClient(5*time.Second, cache)

	c := &config{
		pokeapiClient: client,
	}
	startRepl(c)
}
