package main

import (
	"fmt"
)

func inspect(c *config, key string) error {

	pokemonObj, ok := pokedex[key]

	if !ok {
		fmt.Println("You have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemonObj.Name)
	fmt.Printf("Height: %d\n", pokemonObj.Height)
	fmt.Printf("Weight: %d\n", pokemonObj.Weight)
	fmt.Println("Stats:")

	for i := range pokemonObj.Stats {
		fmt.Printf("  -%s: %d\n", pokemonObj.Stats[i].Stat.Name, pokemonObj.Stats[i].BaseStat)
	}
	fmt.Println("Types:")
	for i := range pokemonObj.Types {
		fmt.Printf("  - %s\n", pokemonObj.Types[i].Type.Name)
	}

	return nil
}
