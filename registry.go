package main

type cliCommand struct {
	name        string
	description string
	cfg         *config
	callback    func(*config, *client, string) error
}

type config struct {
	next     *string
	previous *string
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous map locations",
			callback:    commandPreMap,
		},
		"explore": {
			name:        "explore",
			description: "Display all pokemon present in the selected zone",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon you got",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Inspect the pokemons you caught",
			callback:    commandPokedex,
		},
	}
}
