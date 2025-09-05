package cmd

import (
	"strings"

	"pokedex/internal/client"
)

type pokedexCommand struct {
	caughtPokemon map[string]client.Pokemon
}

func (c pokedexCommand) name() string {
	return "pokedex"
}

func (c pokedexCommand) description() string {
	return "list all caught pokemon"
}

func (c pokedexCommand) run(args []string) (string, error) {
	var b strings.Builder
	b.WriteString("Your Pokedex:\n")
	for pokemonName := range c.caughtPokemon {
		b.WriteString("  - ")
		b.WriteString(pokemonName)
		b.WriteRune('\n')
	}
	return b.String(), nil
}
