package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/maxBRT/pokedex/internal/pokeapi"
	"github.com/maxBRT/pokedex/internal/pokecache"
)

// commandMap handles the "map" command to fetch and display the next 20 locations.
func CommandMap(cfg *pokecache.Config, s string) error {
	// Initialize the NextURL if it's empty.
	if cfg.NextURL == "" {
		cfg.NextURL = BaseURL + "location-area/1/"
	}
	// Extract the current location ID from the URL.
	locationIDstr := strings.TrimPrefix(cfg.NextURL, BaseURL+"location-area/")
	locationIDstr = strings.ReplaceAll(locationIDstr, "/", "")
	locationIDint, err := strconv.Atoi(locationIDstr)
	if err != nil {
		fmt.Println("Error:", err)
	}

	limit := locationIDint + 20 // Fetch the next 20 locations.
	for i := locationIDint; i <= limit; i++ {
		url := BaseURL + "location-area/" + strconv.Itoa(i) + "/"
		cachedLocation, exist := Cache.Get(url) // Check if the location is cached.
		if exist {
			fmt.Println(string(cachedLocation))
			cfg.NextURL = BaseURL + "location-area/" + strconv.Itoa(i+1) + "/"
			cfg.PreviousURL = url
			continue
		}

		cfg.PreviousURL = url
		res, err := http.Get(url) // Fetch the location from the API.
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 { // Skip if the response is not successful.
			limit++
			continue
		}
		decoder := json.NewDecoder(res.Body)
		var location pokeapi.Location
		err = decoder.Decode(&location)
		if err != nil {
			fmt.Println("Error:", err)
		}
		Cache.Add(url, []byte(location.Name)) // Cache the location.
		fmt.Println(location.Name)
		cfg.NextURL = BaseURL + "location-area/" + strconv.Itoa(i+1) + "/"
	}
	return nil
}
