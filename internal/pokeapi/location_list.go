package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const fullURL = baseURL + "/location-area/"

type LocationList struct {
	Count    int              `json:"count"`
	Next     *string          `json:"next"`
	Previous *string          `json:"previous"`
	Results  []LocationResult `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (*LocationList, error) {
	url := fullURL
	if pageURL != nil {
		url = *pageURL
	}
	locationList := LocationList{}

	if body, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(body, &locationList)
		if err != nil {
			return nil, err
		}
		return &locationList, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &locationList)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, body)
	return &locationList, nil
}
