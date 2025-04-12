package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var listOfCommands map[string]cliCommand
var conf config

func start_repl() {
	conf = config{
		next:     "https://pokeapi.co/api/v2/location-area/1",
		previous: "",
	}

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
		"map": {
			name:        "map",
			description: "Displays the 20 next area location",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the 20 previous area location",
			callback:    commandMapb,
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
			err := command.callback(&conf)
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

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
}
