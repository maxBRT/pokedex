package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/maxBRT/pokedex/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Config struct {
	NextURL     string
	PreviousURL string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Location struct {
	Name string `json:"name"`
}

var (
	cliCmdMap map[string]cliCommand
	mapCalls  int
)

func main() {
	cliCmdMap = map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "List all available commands",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Get the map of 20 new locations with each call",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Go back 20 locations on the map",
			Callback:    commandMapBack,
		},
	}

	bufioScanner := bufio.NewScanner(os.Stdin)
	config := &Config{NextURL: "", PreviousURL: ""}
	for {
		fmt.Print("Pokedex > ")
		bufioScanner.Scan()
		input := bufioScanner.Text()
		words := cleanInput(input)
		for _, word := range words {
			cmd, ok := cliCmdMap[word]
			if ok {
				cmd.Callback(config)
				if cmd.Name == "map" {
				}
			}
		}
	}
}

func cleanInput(input string) []string {
	var words []string
	words = strings.Fields(strings.ToLower(input))
	for i := range words {
		fmt.Println(words[i])
	}
	return words
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for _, cmd := range cliCmdMap {
		fmt.Printf("  %s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func commandMap(cfg *Config) error {
	// Use cfg.NextURL to fetch the next page
	if cfg.NextURL == "" {
		cfg.NextURL = baseURL + "location-area/1/"
	}
	locationIDstr := strings.TrimPrefix(cfg.NextURL, baseURL+"location-area/")
	locationIDstr = strings.ReplaceAll(locationIDstr, "/", "")
	locationIDint, err := strconv.Atoi(locationIDstr)
	if err != nil {
		fmt.Println("Error:", err)
	}

	limit := locationIDint + 20
	for i := locationIDint; i <= limit; i++ {
		url := baseURL + "location-area/" + strconv.Itoa(i) + "/"
		cfg.PreviousURL = url
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			limit++
			continue
		}
		decoder := json.NewDecoder(res.Body)
		var location Location
		err = decoder.Decode(&location)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(location.Name)
		cfg.NextURL = baseURL + "location-area/" + strconv.Itoa(i+1) + "/"
	}
	return nil
}

func commandMapBack(cfg *Config) error {
	if cfg.NextURL == "" {
		fmt.Println("You are already at the beginning of the map.")
		return nil
	}
	locationIDstr := strings.TrimPrefix(cfg.PreviousURL, baseURL+"location-area/")
	locationIDstr = strings.ReplaceAll(locationIDstr, "/", "")
	locationIDint, err := strconv.Atoi(locationIDstr)
	if err != nil {
		fmt.Println("Error:", err)
	}

	limit := locationIDint - 20
	for i := locationIDint; i >= limit; i-- {
		url := baseURL + "location-area/" + strconv.Itoa(i) + "/"
		cfg.PreviousURL = url
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			limit++
			continue
		}
		decoder := json.NewDecoder(res.Body)
		var location Location
		err = decoder.Decode(&location)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(location.Name)
		cfg.NextURL = baseURL + "location-area/" + strconv.Itoa(i) + "/"
	}
	return nil
}
