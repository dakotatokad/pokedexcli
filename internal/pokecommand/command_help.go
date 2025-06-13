package pokecommand

import (
	"fmt"

	"github.com/dakotatokad/pokedexcli/internal/pokeconfig"
)

func CallbackHelp(cfg *pokeconfig.Config, args ...string) error {
	println("Welcome to Pokedex help menu!")
	println("Available commands:")

	availableCommands := GetCommands()
	for _, cmd := range availableCommands {
		fmt.Println(cmd.Name + ": " + cmd.Description)
	}

	fmt.Println("")

	return nil
}
