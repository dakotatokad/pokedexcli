package pokecommand

import (
	"errors"
	"fmt"

	"github.com/dakotatokad/pokedexcli/internal/pokeconfig"
)

func callbackMap(cfg *pokeconfig.Config, args ...string) error {

	resp, err := cfg.PokeapiClient.ListLocationAreas(cfg.NextLocationAreaURL)
	if err != nil {
		println("Error fetching location areas:", err.Error())
	}
	fmt.Println("Available Location Areas:")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	cfg.NextLocationAreaURL = resp.Next
	cfg.PreviousLocationAreaURL = resp.Previous

	return nil
}

func callbackMapBack(cfg *pokeconfig.Config, args ...string) error {

	if cfg.PreviousLocationAreaURL == nil {
		return errors.New("you're on the first page, cannot go back")
	}

	resp, err := cfg.PokeapiClient.ListLocationAreas(cfg.PreviousLocationAreaURL)
	if err != nil {
		println("Error fetching location areas:", err.Error())
	}
	fmt.Println("Available Location Areas:")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	cfg.NextLocationAreaURL = resp.Next
	cfg.PreviousLocationAreaURL = resp.Previous

	return nil
}
