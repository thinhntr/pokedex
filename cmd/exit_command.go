package cmd

type exitCommand struct{}

func (c exitCommand) name() string {
	return "exit"
}

func (c exitCommand) description() string {
	return "Exit the Pokedex"
}

func (c exitCommand) run() (string, error) {
	return "Closing the Pokedex... Goodbye!", ErrProgramQuit
}
