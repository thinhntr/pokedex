package main

import (
	"fmt"
	"os"

	"pokedex/cmd"
)

func main() {
	if err := cmd.NewProgram().Run(); err != nil {
		fmt.Fprintf(os.Stderr, "\nencountered error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "bye!")
}
