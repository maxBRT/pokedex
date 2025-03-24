package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maxBRT/pokedex/internal/pokeapi"
	"github.com/maxBRT/pokedex/internal/pokecache"
)

func CommandExplore(cfg *pokecache.Config, locationName string) error {
	var location pokeapi.Location
	url := BaseURL + "location-area/" + locationName + "/"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Location not found.")
		return nil
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&location)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, encounter := range location.Encounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
