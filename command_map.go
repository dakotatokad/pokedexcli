package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
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

func callbackMapBack(cfg *config, args ...string) error {

	if cfg.previousLocationAreaURL == nil {
		return errors.New("you're on the first page, cannot go back")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
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
