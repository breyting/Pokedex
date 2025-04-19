package main

import (
	"fmt"
)

func commandinspect(config *config, input []string) error {
	pokemon := input[0]
	val, ok := ownPokedex[pokemon]
	if ok {
		printInfo(val)
		return nil
	} else {
		return fmt.Errorf("you have not coaught that pokemon")
	}
}

func printInfo(pokemonInfo Pokemon) {
	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, val := range pokemonInfo.Stats {
		fmt.Printf("-%s: %d\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Println("Types:")
	for _, val := range pokemonInfo.Types {
		fmt.Printf("- %s\n", val.Type.Name)
	}
}
