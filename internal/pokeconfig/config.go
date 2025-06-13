package pokeconfig

import "github.com/dakotatokad/pokedexcli/internal/pokeapi"

type Config struct {
	PokeapiClient           pokeapi.Client
	NextLocationAreaURL     *string
	PreviousLocationAreaURL *string
	CaughtPokemon           map[string]pokeapi.Pokemon
}
