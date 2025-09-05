package cmd

import (
	"errors"
	"fmt"
	"strings"

	"pokedex/internal/client"
)

type inspectCommand struct {
	client        *client.Client
	caughtPokemon map[string]client.Pokemon
}

func (c inspectCommand) name() string {
	return "inspect"
}

func (c inspectCommand) description() string {
	return "inspect a pokemon, e.g. `inspect zekrom`"
}

func (c inspectCommand) run(args []string) (string, error) {
	if len(args) < 1 || args[0] == "" {
		return "", errors.New("missing pokemon argument")
	}

	pokemonName := args[0]
	pokemon, err := c.client.GetPokemon(pokemonName)
	if err != nil {
		return "", err
	}

	if _, wasCaught := c.caughtPokemon[pokemonName]; !wasCaught {
		return "you have not caught that pokemon", nil
	}

	var b strings.Builder

	b.WriteString(fmt.Sprintf("Name: %s\n", pokemon.Name))
	b.WriteString(fmt.Sprintf("Height: %d\n", pokemon.Height))
	b.WriteString(fmt.Sprintf("Weight: %d\n", pokemon.Weight))

	b.WriteString("Stats:\n")

	tmpStatsMap := make(map[string]int)
	for _, stat := range pokemon.Stats {
		tmpStatsMap[stat.Stat.Name] = stat.BaseStat
	}
	b.WriteString(fmt.Sprintf("  -hp: %d\n", tmpStatsMap["hp"]))
	b.WriteString(fmt.Sprintf("  -attack: %d\n", tmpStatsMap["attack"]))
	b.WriteString(fmt.Sprintf("  -defense: %d\n", tmpStatsMap["defense"]))
	b.WriteString(fmt.Sprintf("  -special-attack: %d\n", tmpStatsMap["special-attack"]))
	b.WriteString(fmt.Sprintf("  -special-defense: %d\n", tmpStatsMap["special-defense"]))
	b.WriteString(fmt.Sprintf("  -speed: %d\n", tmpStatsMap["speed"]))

	b.WriteString("Types:\n")
	for _, t := range pokemon.Types {
		b.WriteString(fmt.Sprintf("  - %s\n", t.Type.Name))
	}

	return b.String(), nil
}
