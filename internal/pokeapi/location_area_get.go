package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetLocationAreaDetails(areaName string) (*LocationAreaDetail, error) {
	url := baseURL + "/location-area/" + areaName + "/"
	locationAreaDetail := LocationAreaDetail{}

	if body, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(body, &locationAreaDetail)
		if err != nil {
			return nil, err
		}
		return &locationAreaDetail, nil
	}

	req, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &locationAreaDetail)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, body)
	return &locationAreaDetail, nil
}

