package client

import (
	"encoding/json"
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
