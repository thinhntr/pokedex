package cmd

import (
	"strings"

	"pokedex/internal/client"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func initSpecialCommands() []command {
	urls := mapCommandURLs{}
	c := client.NewClient()
	caughtPokemon := make(map[string]client.Pokemon)

	mapCmd := mapCommand{client: c, urls: &urls}
	mapbCmd := mapbCommand{client: c, urls: &urls}
	exploreCmd := exploreCommand{client: c}
	pokedexCmd := pokedexCommand{caughtPokemon: caughtPokemon}
	catchCmd := catchCommand{
		client:        c,
		caughtPokemon: caughtPokemon,
	}
	inspectCmd := inspectCommand{
		client:        c,
		caughtPokemon: caughtPokemon,
	}

	return []command{mapCmd,
		mapbCmd,
		exploreCmd,
		catchCmd,
		inspectCmd,
		pokedexCmd,
	}
}
