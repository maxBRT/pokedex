package main

import (
	"github.com/maxBRT/pokedex/internal/pokecache"
)

// cliCommand represents a CLI command with its name, description, and callback function.
type CliCommand struct {
	Name        string
	Description string
	Callback    func(*pokecache.Config, string) error
}
