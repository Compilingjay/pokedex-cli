package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    cmdHelp,
		},
		"map": {
			name:        "mapf",
			description: "Displays the name of the next 20 locations",
			callback:    cmdMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the name of the previous 20 locations",
			callback:    cmdMapb,
		},
		"explore": {
			name:        "explore <location name>",
			description: "Displays the list of pokemon in a location",
			callback:    cmdExplore,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "Attempt to catch a pokemon by their name",
			callback:    cmdCatch,
		},
		"inspect": {
			name:        "inspect <pokemon name>",
			description: "Look at a pokemon that you have already caught",
			callback:    cmdInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all the pokemon that you have caught",
			callback:    cmdPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    cmdExit,
		},
	}
}
