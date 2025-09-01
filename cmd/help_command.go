package cmd

import (
	"strings"
)

type helpCommand struct {
	commands []command
}

func (c helpCommand) name() string {
	return "help"
}

func (c helpCommand) description() string {
	return "displays a help message"
}

func (c helpCommand) run(args []string) (string, error) {
	var b strings.Builder

	b.WriteString("Welcome to the Pokedex!\n")
	b.WriteString("Usage:\n\n")

	for _, cmd := range c.commands {
		if cmd != nil {
			b.WriteString(cmd.name())
			b.WriteString(": ")
			b.WriteString(cmd.description())
			b.WriteRune('\n')
		}
	}

	return b.String(), nil
}
