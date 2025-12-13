package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//pokeCache "BdPokedex/internal/cache"
)

func startRepl() {
	//fmt.Println("===Pokedex===")
	reader := bufio.NewScanner(os.Stdin)
	commands := GetCommands()
	cfg := &config{}
	client := NewClient()

	for {
		fmt.Print("Pokedex >")
		reader.Scan()

		words := cleanInput((reader.Text()))
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		argument00 := ""

		if len(words) > 1 {
			argument00 = words[1]
		}

		command, exists := commands[commandName]
		if exists {
			err := command.callback(cfg, client, argument00)
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
