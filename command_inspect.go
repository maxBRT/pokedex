package main

import (
	"fmt"

	"github.com/maxBRT/pokedex/internal/pokecache"
)

func CommandInspect(cfg *pokecache.Config, pokemonName string) error {
	pokemon, ok := UserPokedex[pokemonName]
	if !ok {
		fmt.Println("Pokemon not found in your Pokedex.")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("ID: %d\n", pokemon.ID)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types: ")
	for i, t := range pokemon.Types {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(t.Type.Name)
	}
	fmt.Println()
	return nil
}
