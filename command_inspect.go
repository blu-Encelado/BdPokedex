package main

import (
	"fmt"
)

func commandInspect(cfg *config, c *client, s string) error {
	if s == "" {
		return fmt.Errorf("miss a pokemon to inspect")
	}

	pokemon, ok := c.pokedex[s]
	if ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pokType := range pokemon.Types {
			fmt.Printf("  - %s\n", pokType.Type.Name)
		}
		return nil
	}

	fmt.Printf("You don't have a %s to inspect\n", s)

	return nil
}
