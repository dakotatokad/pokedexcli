package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]

		switch command {
		case "help":
			fmt.Println("Welcome to Pokedex help menu!")
			fmt.Println("Available commands:")
			fmt.Println("  help - Show this help menu")
			fmt.Println("  add <pokemon> - Add a Pokémon to your Pokedex")
			fmt.Println("  list - List all Pokémon in your Pokedex")
			fmt.Println("  remove <pokemon> - Remove a Pokémon from your Pokedex")
			fmt.Println("  exit - Exit the Pokedex REPL")
			case "exit":
				os.Exit(0)
			default:
				fmt.Println("Invalid command: ", command)
		}

		}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
