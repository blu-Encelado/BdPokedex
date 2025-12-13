package main

import (
	"fmt"
)

func commandHelp(cfg *config, client *client, s string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
