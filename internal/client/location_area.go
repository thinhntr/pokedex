package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type LocationAreaDetails struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationArea(pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	rawBytes, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationArea{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}
		defer res.Body.Close()

		rawBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, err
		}
		c.cache.Add(url, rawBytes)
	}

	var locationArea LocationArea
	err := json.Unmarshal(rawBytes, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}

func (c *Client) GetLocationAreaDetails(locationName string) (LocationAreaDetails, error) {
	url := baseURL + "/location-area/" + locationName

	// we should sanitize locationName

	rawBytes, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaDetails{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaDetails{}, err
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return LocationAreaDetails{}, fmt.Errorf("location area \"%s\" not found\n", locationName)
		}

		rawBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaDetails{}, err
		}
		c.cache.Add(url, rawBytes)
	}

	var locationAreaDetails LocationAreaDetails
	err := json.Unmarshal(rawBytes, &locationAreaDetails)
	if err != nil {
		return LocationAreaDetails{}, err
	}
	return locationAreaDetails, nil
}
