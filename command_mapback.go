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

// commandMapBack handles the "mapb" command to fetch and display the previous 20 locations.
func CommandMapBack(cfg *pokecache.Config, s string) error {
	if cfg.NextURL == "" {
		fmt.Println("You are already at the beginning of the map.")
		return nil
	}
	// Extract the current location ID from the URL.
	locationIDstr := strings.TrimPrefix(cfg.PreviousURL, BaseURL+"location-area/")
	locationIDstr = strings.ReplaceAll(locationIDstr, "/", "")
	locationIDint, err := strconv.Atoi(locationIDstr)
	if err != nil {
		fmt.Println("Error:", err)
	}

	limit := locationIDint - 20 // Fetch the previous 20 locations.
	for i := locationIDint; i >= limit; i-- {
		url := BaseURL + "location-area/" + strconv.Itoa(i) + "/"
		cachedLocation, exist := Cache.Get(url) // Check if the location is cached.
		if exist {
			fmt.Println(string(cachedLocation))
			cfg.NextURL = BaseURL + "location-area/" + strconv.Itoa(i-1) + "/"
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
		fmt.Println(location.Name)
		cfg.NextURL = BaseURL + "location-area/" + strconv.Itoa(i) + "/"
	}
	return nil
}
