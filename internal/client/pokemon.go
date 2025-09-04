package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Name           string `json:"name"`
	Order          int    `json:"order"`
	BaseExperience int    `json:"base_experience"`
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	rawBytes, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return Pokemon{}, fmt.Errorf("pokemon \"%s\" not found\n", pokemonName)
		}

		rawBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

		c.cache.Add(url, rawBytes)
	}

	var pokemon Pokemon
	err := json.Unmarshal(rawBytes, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
