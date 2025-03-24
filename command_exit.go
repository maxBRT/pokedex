package main

import (
	"fmt"
	"os"

	"github.com/maxBRT/pokedex/internal/pokecache"
)

// commandExit handles the "exit" command to terminate the program.
func CommandExit(cfg *pokecache.Config, s string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
