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
	client := client.NewClient()
	mapCmd := mapCommand{client: client, urls: &urls}
	mapbCmd := mapbCommand{client: client, urls: &urls}
	exploreCmd := exploreCommand{client: client}
	catchCmd := catchCommand{client: client}
	return []command{mapCmd, mapbCmd, exploreCmd, catchCmd}
}
