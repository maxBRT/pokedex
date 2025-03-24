package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/maxBRT/pokedex/internal/pokeapi"
	"github.com/maxBRT/pokedex/internal/pokecache"
)

func CommandCatch(cfg *pokecache.Config, pokemonName string) error {
	url := BaseURL + "pokemon/" + pokemonName + "/"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("Pokemon not found.")
		return nil
	}
	decoder := json.NewDecoder(res.Body)
	var pokemon pokeapi.Pokemon
	err = decoder.Decode(&pokemon)
	if err != nil {
		fmt.Println("Error:", err)
	}
	catchChance := pokemon.BaseExperience / 5
	if catchChance > 100 {
		catchChance = 90
	}
	catchTry := rand.Intn(100)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if catchTry > catchChance {
		fmt.Printf("Congratulations! You caught %s!\n", pokemon.Name)
		UserPokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s ran away!\n", pokemon.Name)
	}
	return nil
}
