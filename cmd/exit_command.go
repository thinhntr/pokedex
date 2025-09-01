package cmd

type exitCommand struct{}

func (c exitCommand) name() string {
	return "exit"
}

func (c exitCommand) description() string {
	return "exit the Pokedex"
}

func (c exitCommand) run(args []string) (string, error) {
	return "Closing the Pokedex... Goodbye!", ErrProgramQuit
}
