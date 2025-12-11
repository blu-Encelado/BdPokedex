package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("===Pokedex===")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		fmt.Print("Your command was: ", cleanInput(scanner.Text())[0], "\n")

	}

}
