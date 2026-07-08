package main

import (
	"errors"
	"testing"
)

func testExit(t *testing.T) {
	testCases := []struct {
		input    string
		expected error
	}{
		{
			input:    "exit",
			expected: nil,
		},
		{
			input:    "help",
			expected: errors.New("wrong input"),
		},
	}

	for _, c := range testCases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]

			if word != "exit" {
				t.Errorf("wrong input")
			}
		}
	}
}

func testHelp(t *testing.T) {

}

func testMap(t *testing.T, c *config) {

}

func testMapb(t *testing.T) {

}
