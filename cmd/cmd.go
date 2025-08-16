package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var ErrProgramQuit = errors.New("program received signal to quit")

type command interface {
	name() string
	description() string
	run() (string, error)
}

type program struct {
	commands map[string]command
	scanner  *bufio.Scanner
}

func (p *program) Run() error {
	for {
		cmd, err := p.parseCommand()
		if err != nil {
			return err
		}
		if cmd == nil {
			fmt.Println("Unknown command")
			continue
		}

		res, err := cmd.run()
		switch err {
		case nil:
			fmt.Println(res)
		case ErrProgramQuit:
			fmt.Println(res)
			return nil
		default:
			fmt.Fprintf(os.Stdin, "%v\n", err)
		}
	}
}

func (p *program) parseCommand() (command, error) {
	fmt.Print("Pokedex > ")
	if p.scanner.Scan() {
		line := p.scanner.Text()
		tokens := cleanInput(line)
		if len(tokens) < 1 {
			return nil, nil
		}

		cmd, ok := p.commands[tokens[0]]
		if !ok {
			return nil, nil
		}
		return cmd, nil
	}
	if err := p.scanner.Err(); err != nil {
		return nil, err
	}
	return nil, errors.New("program reached EOF")
}

func NewProgram() *program {
	commands := []command{
		helpCommand{},
		exitCommand{},
	}
	helpCmd := helpCommand{
		commands: commands,
	}

	commandMap := make(map[string]command)
	for _, cmd := range commands {
		commandMap[cmd.name()] = cmd
	}
	commandMap[helpCmd.name()] = helpCmd

	return &program{
		commands: commandMap,
		scanner:  bufio.NewScanner(os.Stdin),
	}
}
