package cmd

import (
	"strings"

	"pokedex/internal/client"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func getMapCommands() (mapCommand, mapbCommand) {
	urls := mapCommandURLs{}
	client := client.NewClient()
	mapCmd := mapCommand{client: client, urls: &urls}
	mapbCmd := mapbCommand{client: client, urls: &urls}
	return mapCmd, mapbCmd
}
