package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
)

const (
	decayScalar = 1.0 / 160.0
)

func cmdHelp(cfg *config, _ ...string) error {
	fmt.Println("Usage:")
	fmt.Println("--------------------")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func cmdMapf(cfg *config, _ ...string) error {
	respLocations, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = respLocations.Next
	cfg.prevLocationsURL = respLocations.Previous

	for _, locations := range respLocations.Results {
		fmt.Println(locations.Name)
	}

	return nil
}

func cmdMapb(cfg *config, _ ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page, can't go to previous page")
	}

	respLocations, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = respLocations.Next
	cfg.prevLocationsURL = respLocations.Previous

	for _, locations := range respLocations.Results {
		fmt.Println(locations.Name)
	}

	return nil
}

func cmdExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a location to explore")
	}

	locationName := args[0]

	respLocationArea, err := cfg.pokeapiClient.GetLocationAreaPokemon(&locationName)
	if err != nil {
		return fmt.Errorf("please provide a valid location")
	}

	fmt.Printf("Exploring %s\n", locationName)
	fmt.Println("Found pokemon:")
	for _, encounter := range respLocationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func cmdCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a valid name for a pokemon")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(&name)
	if err != nil {
		return fmt.Errorf("please provide a valid name for a pokemon")
	}

	roll := rand.Float64()
	fmt.Printf("throwing a pokeball at %s...\n", name)

	success := (roll < (math.Pow(0.5, float64(pokemon.BaseExperience)*decayScalar)))
	if !success {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	cfg.caughtPokemon[name] = pokemon
	fmt.Printf("%s was caught!\n", name)

	return nil
}

func cmdInspect(cfg *config, args ...string) error {
	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("  - %s\n", typ.Type.Name)
	}

	return nil
}

func cmdPokedex(cfg *config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}

func cmdExit(cfg *config, _ ...string) error {
	os.Exit(0)
	return nil
}
