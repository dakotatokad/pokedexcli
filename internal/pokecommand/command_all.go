package pokecommand

import "github.com/dakotatokad/pokedexcli/internal/pokeconfig"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*pokeconfig.Config, ...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Show this help menu",
			Callback:    CallbackHelp,
		},
		"map": {
			Name:        "map",
			Description: "List the next page of location areas",
			Callback:    callbackMap,
		},
		"map-back": {
			Name:        "map-back",
			Description: "Go back to the previous page of location areas",
			Callback:    callbackMapBack,
		},
		"explore": {
			Name:        "explore {location_area_name}",
			Description: "Explore a specific location area by name for Pokemon encounters",
			Callback:    callbackExplore,
		},
		"catch": {
			Name:        "catch {pokemon_name}",
			Description: "Catch a Pokemon by name",
			Callback:    CallbackCatch,
		},
		"inspect": {
			Name:        "inspect {pokemon_name}",
			Description: "Inspect a caught Pokemon by name",
			Callback:    callbackInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Show the Pokedex with caught Pokemon",
			Callback:    callbackPokedex,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex REPL",
			Callback:    callbackExit,
		},
	}
}
