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
	run(args []string) (string, error)
}

type program struct {
	commands map[string]command
	scanner  *bufio.Scanner
}

func (p *program) Run() error {
	for {
		cmd, args, err := p.parseCommand()
		if err != nil {
			return err
		}
		if cmd == nil {
			fmt.Println("Unknown command")
			continue
		}

		res, err := cmd.run(args)
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

func (p *program) parseCommand() (cmd command, args []string, err error) {
	fmt.Print("Pokedex > ")
	if p.scanner.Scan() {
		line := p.scanner.Text()
		tokens := cleanInput(line)
		if len(tokens) < 1 {
			return nil, []string{}, nil
		}

		cmd, ok := p.commands[tokens[0]]
		if !ok {
			return nil, []string{}, nil
		}
		return cmd, tokens[1:], nil
	}
	if err := p.scanner.Err(); err != nil {
		return nil, []string{}, err
	}
	return nil, []string{}, errors.New("program reached EOF")
}

func NewProgram() *program {
	commands := []command{
		helpCommand{},
		exitCommand{},
	}
	commands = append(commands, initSpecialCommands()...)
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
