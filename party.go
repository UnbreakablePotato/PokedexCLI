package main

import (
	"fmt"
	"sync"
)

var mu sync.RWMutex

var pokemonParty = make(map[string]pokemon)

func addPokemon(pokemon pokemon, party map[string]pokemon) error {

	if len(party) > 6 {
		fmt.Printf("Your party is full!\nCould not add %s\n", pokemon.Name)
		return nil
	}

	mu.Lock()

	defer mu.Unlock()

	pname := pokemon.Name

	party[pname] = pokemon
	fmt.Printf("%s was added to your party!\n", pokemon.Name)

	return nil
}

func removePokemon(pokemon string, party map[string]pokemon) error {
	mu.Lock()

	defer mu.Unlock()

	delete(party, pokemon)
	fmt.Printf("%s has been removed from your party!\n", pokemon)
	return nil
}
