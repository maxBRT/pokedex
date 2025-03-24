package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/maxBRT/pokedex/internal/pokecache"
)

func InitRepl() {
	CliCommandsMap = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "List all available commands",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Get the map of 20 new locations with each call",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Go back 20 locations on the map",
			Callback:    CommandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore the location and its encounters. Type 'explore <location-name>'",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon. Type 'catch <pokemon-name>'",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a pokemon in your Pokedex. Type 'inspect <pokemon-name>'",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List all pokemon in your Pokedex",
			Callback:    CommandPokedex,
		},
	}

	bufioScanner := bufio.NewScanner(os.Stdin) // Scanner for reading user input.
	config := &pokecache.Config{NextURL: "", PreviousURL: ""}
	for {
		fmt.Print("Pokedex > ") // Prompt for user input.
		bufioScanner.Scan()
		input := bufioScanner.Text()
		words := CleanInput(input) // Clean and parse the input.
		for _, word := range words {
			cmd, ok := CliCommandsMap[word]
			if ok {
				switch word {
				case "pokedex":
					if len(UserPokedex) == 0 {
						fmt.Println("Your Pokedex is empty.")
						continue
					}
					cmd.Callback(config, " ")
					continue
				case "exit":
					cmd.Callback(config, " ")
					return
				case "help":
					cmd.Callback(config, " ")
					continue
				case "map":
					cmd.Callback(config, " ")
					continue
				case "mapb":
					cmd.Callback(config, " ")
					continue
				case "explore":
					if len(words) < 2 {
						fmt.Println("Please provide a location name.")
						continue
					}
					cmd.Callback(config, words[1])
					continue
				case "catch":
					if len(words) < 2 {
						fmt.Println("Please provide a pokemon name.")
						continue
					}
					cmd.Callback(config, words[1])
					continue
				case "inspect":
					if len(words) < 2 {
						fmt.Println("Please provide a pokemon name.")
						continue
					}
					cmd.Callback(config, words[1])
					continue
				default:
					fmt.Println("Invalid command. Type 'help' to see available commands.")
				}
			}
		}
	}
}
