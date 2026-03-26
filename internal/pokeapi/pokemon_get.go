package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name + "/"
	pokemon := Pokemon{}

	if body, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(body, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)
	return pokemon, nil
}
