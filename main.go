package main

import (
	"fmt"

	"github.com/dakotatokad/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		println("Error fetching location areas:", err.Error())
	}
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	//startRepl()
}
