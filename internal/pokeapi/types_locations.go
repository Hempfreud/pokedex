package pokeapi

type LocationResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaDetail struct {
	PokemonEncounters []struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
} `json:"pokemon_encounters"`
}