package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("catch command requires exactly one argument: the name of the pokemon")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("you haven't caught %s yet", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("  Height: %d\n", pokemon.Height)
	fmt.Printf("  Weight: %d\n", pokemon.Weight)
	fmt.Printf("  Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf(" - %s (Hidden: %t)\n", ability.Ability.Name, ability.IsHidden)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
