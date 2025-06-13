package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("catch command requires exactly one argument: the name of the pokemon")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		println("Error fetching pokemon:", err.Error())
	}

	const threshold = 50 // Example threshold for catching a Pokemon
	randNumber := rand.Intn(pokemon.BaseExperience)
	log.Printf("Attempting to catch %s with a random number: %d (threshold: %d)\n", pokemonName, randNumber, threshold)
	if randNumber < threshold {
		fmt.Printf("You caught a %s!\n", pokemonName)
	} else {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	return nil
}
