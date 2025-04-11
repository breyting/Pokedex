package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
