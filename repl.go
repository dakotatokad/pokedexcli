package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command:", commandName)
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show this help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "List the next page of location areas",
			callback:    callbackMap,
		},
		"map-back": {
			name:        "map-back",
			description: "Go back to the previous page of location areas",
			callback:    callbackMapBack,
		},
		"explore": {
			name:        "explore {location_area_name}",
			description: "Explore a specific location area by name for Pokemon encounters",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Catch a Pokemon by name",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Inspect a caught Pokemon by name",
			callback:    callbackInspect,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex REPL",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
