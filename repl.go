package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	//fmt.Println("===Pokedex===")
	reader := bufio.NewScanner(os.Stdin)
	commands := GetCommands()

	for {
		fmt.Print("Pokedex >")
		reader.Scan()

		words := cleanInput((reader.Text()))
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknow command")
		}

	}
}
func cleanInput(text string) []string {
	newText := strings.Fields(strings.TrimSpace(strings.ToLower(text)))

	return newText
}
