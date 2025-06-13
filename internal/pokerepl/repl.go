package pokerepl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dakotatokad/pokedexcli/internal/pokecommand"
	"github.com/dakotatokad/pokedexcli/internal/pokeconfig"
)

func StartRepl(cfg *pokeconfig.Config) {
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

		availableCommands := pokecommand.GetCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command:", commandName)
			continue
		}

		err := command.Callback(cfg, args...)
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
