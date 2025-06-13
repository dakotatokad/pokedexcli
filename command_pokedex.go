package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("you haven't caught any Pokémon yet")
	}

	numberOfCaught := len(cfg.caughtPokemon)

	fmt.Printf("You've caught %d Pokémon:", numberOfCaught)
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s (Base Experience: %d)\n",
			pokemon.Name, pokemon.BaseExperience)
	}

	return nil
}
