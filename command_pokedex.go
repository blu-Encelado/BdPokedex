package main

import (
	"fmt"
)

func commandPokedex(cfg *config, c *client, s string) error {

	if len(c.pokedex) == 0 {
		return fmt.Errorf("your pokedex is empty")
	}
	fmt.Println("Your Pokedex:")
	for name := range c.pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
