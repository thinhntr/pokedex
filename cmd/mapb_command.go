package cmd

import (
	"errors"
	"strings"

	"pokedex/internal/client"
)

type mapbCommand struct {
	client *client.Client
	urls   *mapCommandURLs
}

func (c mapbCommand) name() string {
	return "mapb"
}

func (c mapbCommand) description() string {
	return "displays the previous 20 location area names in the Pokemon world"
}

func (c mapbCommand) run() (string, error) {
	if c.urls.Previous == nil {
		return "", errors.New("first page of location area")
	}
	res, err := c.client.GetLocationArea(c.urls.Previous)
	if err != nil {
		return "", err
	}

	c.urls.Next = res.Next
	c.urls.Previous = res.Previous

	var b strings.Builder

	for _, item := range res.Results {
		b.WriteString(item.Name)
		b.WriteRune('\n')
	}

	return b.String(), nil
}
