package main

import (
	"time"

	pokecache "github.com/maxBRT/pokedex/internal"
)

const baseURL = "https://pokeapi.co/api/v2/"

// Config holds the URLs for navigating through the location map.
type Config struct {
	NextURL     string
	PreviousURL string
}

// cliCommand represents a CLI command with its name, description, and callback function.
type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, string) error
}

// Location represents a location fetched from the API.
type Location struct {
	Name       string `json:"name"`
	Encounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

var CliCommandsMap = make(map[string]cliCommand) // Map of CLI commands.

var c = pokecache.NewCache(10 * time.Second) // Initialize the cache with a 10-second expiration.

var UserPokedex = make(map[string]Pokemon)

func main() {
	InitRepl(&CliCommandsMap)
}
