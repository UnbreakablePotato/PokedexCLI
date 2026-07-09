package main

import (
	"fmt"
)

func printPokedex(c *config) error {
	fmt.Println("Your Pokedex:")
	for i := range pokedex {
		fmt.Printf(" - %s\n", i)
	}

	return nil
}
