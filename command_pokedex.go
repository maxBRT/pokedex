package main

import (
	"fmt"

	"github.com/maxBRT/pokedex/internal/pokecache"
)

func CommandPokedex(cfg *pokecache.Config, s string) error {
	fmt.Println("Your Pokedex:")
	for name := range UserPokedex {
		fmt.Println("- " + name)
	}
	return nil
}
