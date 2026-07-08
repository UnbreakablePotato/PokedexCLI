package main

import (
	"fmt"
	"testing"
)

func TestClean(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ChariZARD, PIKAchu charmander ",
			expected: []string{"charizard", "pikachu", "charmander"},
		},
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " w e r ",
			expected: []string{"w", "e", "r"},
		},
		{
			input:    " WHA T DO YOU M E A N    ",
			expected: []string{"wha", "t", "do", "you", "m", "e", "a", "n"},
		},
	}

	for _, c := range testCases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of the expected and actual slice of word do not match")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				fmt.Printf("Expected: %s | Actual: %s\n", expectedWord, word)
				t.Errorf("The words do not match")
			}
		}
	}
}
