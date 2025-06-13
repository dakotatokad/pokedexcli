package pokecommand

import (
	"os"

	"github.com/dakotatokad/pokedexcli/internal/pokeconfig"
)

func callbackExit(cfg *pokeconfig.Config, args ...string) error {
	println("Exiting Pokedex REPL. Goodbye!")
	os.Exit(0)
	return nil
}
