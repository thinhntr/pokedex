package cmd

import (
	"strings"

	"pokedex/internal/client"
)

type mapCommandURLs struct {
	Next     *string
	Previous *string
}

type mapCommand struct {
	client *client.Client
	urls   *mapCommandURLs
}

func (c mapCommand) name() string {
	return "map"
}

func (c mapCommand) description() string {
	return "displays the next 20 location area names in the Pokemon world"
}

func (c mapCommand) run() (string, error) {
	res, err := c.client.GetLocationArea(c.urls.Next)
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
