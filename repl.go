package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func start_repl() {
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scan.Scan()
		command := scan.Text()
		cleanCommand := cleanInput(command)

		if len(cleanCommand) == 0 {
			continue
		}
		fmt.Printf("Your command was: %s\n", cleanCommand[0])
	}
}
