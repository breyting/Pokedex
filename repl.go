package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var listOfCommands map[string]cliCommand

func start_repl() {
	listOfCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scan.Scan()
		text := scan.Text()
		input := cleanInput(text)

		if len(input) == 0 {
			continue
		}

		command, exists := listOfCommands[input[0]]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage :\n")

	for _, command := range listOfCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
