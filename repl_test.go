package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Charmander Bulbasaur PIKACHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			// Add a new test cases
			input:    "  hello  gopher  ",
			expected: []string{"hello", "gopher"},
		},
		{
			// Add a new test cases
			input:    "  hello  gopher  world  ",
			expected: []string{"hello", "gopher", "world"},
		},
		{ // Add a new test cases
			input:    "  hello  gopher  world  pokemon  ",
			expected: []string{"hello", "gopher", "world", "pokemon"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(
			actual,
		) != len(
			c.expected,
		) { // Check the length of the actual slice against the expected slice
			t.Errorf("Received: %v, Expected : %v", len(actual), len(c.expected))
		}
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("Received: %v, Expected : %v", word, expectedWord)
			}
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
