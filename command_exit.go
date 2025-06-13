package main

import "os"

func callbackExit(cfg *config, args ...string) error {
	println("Exiting Pokedex REPL. Goodbye!")
	os.Exit(0)
	return nil
}
