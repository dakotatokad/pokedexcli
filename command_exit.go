package main

import "os"

func callbackExit() {
	println("Exiting Pokedex REPL. Goodbye!")
	os.Exit(0)
}
