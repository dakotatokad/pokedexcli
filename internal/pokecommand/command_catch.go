package pokecommand

import (
	"errors"
	"fmt"
	"log"
	"math/rand"

	"github.com/dakotatokad/pokedexcli/internal/pokeconfig"
)

func CallbackCatch(cfg *pokeconfig.Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("catch command requires exactly one argument: the name of the pokemon")
	}
	pokemonName := args[0]

	pokemon, err := cfg.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		println("Error fetching pokemon:", err.Error())
	}

	const threshold = 50 // Example threshold for catching a Pokemon
	randNumber := rand.Intn(pokemon.BaseExperience)
	log.Printf("Attempting to catch %s with a random number: %d (threshold: %d)\n", pokemonName, randNumber, threshold)
	if randNumber < threshold {
		fmt.Printf("You caught a %s!\n", pokemonName)
		cfg.CaughtPokemon[pokemonName] = pokemon
	} else {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	return nil
}
