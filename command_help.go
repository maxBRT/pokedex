package main

import (
	"fmt"

	"github.com/maxBRT/pokedex/internal/pokecache"
)

// commandHelp handles the "help" command to display available commands.
func CommandHelp(cfg *pokecache.Config, s string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for _, cmd := range CliCommandsMap {
		fmt.Printf("  %s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
