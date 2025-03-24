// Description: This file contains the Location struct which is used to store the data of a location.
package pokeapi

type Location struct {
	Name       string `json:"name"`
	Encounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
