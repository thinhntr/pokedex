package cmd

import (
	"errors"
	"strings"

	"pokedex/internal/client"
)

type exploreCommand struct {
	client *client.Client
}

func (c exploreCommand) name() string {
	return "explore"
}

func (c exploreCommand) description() string {
	return "explore pokemons in a location area, e.g. `explore mt-coronet-6f`"
}

func (c exploreCommand) run(args []string) (string, error) {
	if len(args) < 1 || args[0] == "" {
		return "", errors.New("missing location area argument")
	}

	area := args[0]
	res, err := c.client.GetLocationAreaDetails(area)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	b.WriteString("Exploring ")
	b.WriteString(area)
	b.WriteString("...\n")

	for _, item := range res.PokemonEncounters {
		b.WriteString("- ")
		b.WriteString(item.Pokemon.Name)
		b.WriteRune('\n')
	}

	return b.String(), nil
}
