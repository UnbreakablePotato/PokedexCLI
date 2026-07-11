package main

import "testing"

var writetestMap = make(map[string]pokemon)

func TestWrite(t *testing.T) {
	testCases := []struct {
		input    map[string]pokemon
		expected error
	}{
		{
			input:    writetestMap,
			expected: nil,
		},
	}
	for _, c := range testCases {
		if err := writeParty(c.input); err != nil {
			t.Errorf("Failed to write file...")
		}
	}
}
