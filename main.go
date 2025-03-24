package main

import (
	"time"

	"github.com/maxBRT/pokedex/internal/pokeapi"
	"github.com/maxBRT/pokedex/internal/pokecache"
)

const BaseURL = "https://pokeapi.co/api/v2/"

var CliCommandsMap map[string]CliCommand // Map of available commands.

var Cache = pokecache.NewCache(
	10 * time.Second,
) // Initialize the cache with a 10-second expiration.

var UserPokedex = make(map[string]pokeapi.Pokemon) // Map of caught Pokemon.

func main() {
	InitRepl()
}
