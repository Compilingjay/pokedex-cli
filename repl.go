package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func replStart(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	commandList := getCommands()

	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()

		args := parseInput(input)
		cmd, exists := commandList[args[0]]
		if exists {
			err := cmd.callback(cfg, args[1:]...)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Error: invalid command.")
		}
	}
}

func parseInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
