package main

import (
	"time"

	"github.com/dakotatokad/pokedexcli/internal/pokeapi"
	"github.com/dakotatokad/pokedexcli/internal/pokeconfig"
	"github.com/dakotatokad/pokedexcli/internal/pokerepl"
)

func main() {

	cfg := pokeconfig.Config{
		PokeapiClient:           pokeapi.NewClient(time.Hour),
		NextLocationAreaURL:     nil,
		PreviousLocationAreaURL: nil,
		CaughtPokemon:           make(map[string]pokeapi.Pokemon),
	}

	pokerepl.StartRepl(&cfg)
}
