package main

import "fmt"

func callbackHelp() {
	println("Welcome to Pokedex help menu!")
	println("Available commands:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Println(cmd.name + ": " + cmd.description)
	}

	fmt.Println("")
}
