package main

import "fmt"

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage :\n")

	for _, command := range listOfCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
