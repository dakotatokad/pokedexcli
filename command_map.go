package main

import (
	"fmt"

	"github.com/dakotatokad/pokedexcli/internal/pokeapi"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		println("Error fetching location areas:", err.Error())
	}
	fmt.Println("Available Location Areas:")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	return nil
}
