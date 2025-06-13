package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("explore command requires exactly one argument: the name of the location area")
	}
	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.GetLocationAreas(locationAreaName)
	if err != nil {
		println("Error fetching location areas:", err.Error())
	}
	fmt.Println("Location Area Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
