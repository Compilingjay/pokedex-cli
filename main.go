package main

import (
	"time"

	"pokeapi"
)

func main() {
	client := pokeapi.NewClient(10*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: client,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	replStart(cfg)
}
