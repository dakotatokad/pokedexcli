package main

import "os"

func callbackExit() error {
	println("Exiting Pokedex REPL. Goodbye!")
	os.Exit(0)
	return nil
}
