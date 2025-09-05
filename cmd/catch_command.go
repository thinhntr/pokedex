package cmd

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"

	"pokedex/internal/client"
)

type catchCommand struct {
	client *client.Client
}

func (c catchCommand) name() string {
	return "catch"
}

func (c catchCommand) description() string {
	return "catch a pokemon, e.g. `catch zekrom`"
}

func (c catchCommand) run(args []string) (string, error) {
	if len(args) < 1 || args[0] == "" {
		return "", errors.New("missing pokemon argument")
	}

	pokemon := args[0]
	res, err := c.client.GetPokemon(pokemon)
	if err != nil {
		return "", err
	}

	var builder strings.Builder

	e := float64(res.BaseExperience)
	catchScore := math.Abs(math.Pow(e, 0.9))
	rollScore := rand.Float64() * e
	wasCaught := rollScore < catchScore

	builder.WriteString(fmt.Sprintf("Throwing a Pokeball at %s...\n", pokemon))

	if wasCaught {
		builder.WriteString(fmt.Sprintf("%s was caught!\n", pokemon))
	} else {
		builder.WriteString(fmt.Sprintf("you rolled: %.3f\n", rollScore))
		builder.WriteString(fmt.Sprintf("%s escaped!\n", pokemon))
	}

	return builder.String(), nil

}
