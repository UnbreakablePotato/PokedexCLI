package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func writeParty(party map[string]pokemon) error {
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

func readParty(filename string) error {

	return nil
}
