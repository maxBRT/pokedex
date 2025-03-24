package main

import (
	"strings"
)

// cleanInput processes the user input by converting it to lowercase and splitting it into words.
func CleanInput(input string) []string {
	var words []string
	words = strings.Fields(strings.ToLower(input))
	return words
}
