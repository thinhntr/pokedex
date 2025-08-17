package client

import (
	"encoding/json"
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

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	var locationArea LocationArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}
