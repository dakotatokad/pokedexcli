package pokecommand

import (
	"fmt"

	"github.com/dakotatokad/pokedexcli/internal/pokeconfig"
)

func callbackPokedex(cfg *pokeconfig.Config, args ...string) error {

	if len(cfg.CaughtPokemon) == 0 {
		return fmt.Errorf("you haven't caught any Pokémon yet")
	}

	numberOfCaught := len(cfg.CaughtPokemon)

	fmt.Printf("You've caught %d Pokémon:\n", numberOfCaught)
	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Printf("- %s (Base Experience: %d)\n",
			pokemon.Name, pokemon.BaseExperience)
	}

	return nil
}
