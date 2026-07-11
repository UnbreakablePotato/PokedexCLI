package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func writeParty(party map[string]pokemon) error {
	//path := filepath.Join(os.TempDir(), "party")
	file, err := os.Create("party")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return err
	}
	defer file.Close()

	writable, err := json.Marshal(party)
	if err != nil {
		fmt.Println("Could not marshal party map")
		return err
	}

	bytes, err := file.Write(writable)

	fmt.Printf("Wrote %d bytes\n", bytes)

	return nil
}

func loadParty(filename string) error {
	//path := filepath.Join(os.TempDir(), filename)

	fileData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not open file: %v\n", err)
		return err
	}

	if err := json.Unmarshal(fileData, &pokemonParty); err != nil {
		fmt.Printf("Failed to unmarshal file into pokemon party map: %v", err)
		return err
	}

	return nil
}
