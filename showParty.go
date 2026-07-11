package main

import "fmt"

func commandShowParty(c *config) error {

	if len(pokemonParty) < 1 {
		fmt.Println("No pokemon in active roster...")
		return nil
	}

	for k := range pokemonParty {
		fmt.Printf("%s\n", k)
	}
	return nil
}
