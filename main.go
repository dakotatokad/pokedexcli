package main

import "github.com/dakotatokad/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {

	cfg := config{
		pokeapiClient:           pokeapi.NewClient(),
		nextLocationAreaURL:     nil,
		previousLocationAreaURL: nil,
	}

	startRepl(&cfg)
}
