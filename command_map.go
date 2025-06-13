package main

import (
	"fmt"
)

func callbackMap(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas()
	if err != nil {
		println("Error fetching location areas:", err.Error())
	}
	fmt.Println("Available Location Areas:")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous

	return nil
}
