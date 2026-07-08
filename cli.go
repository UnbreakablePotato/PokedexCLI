package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type config struct {
	URL     string
	Count   int      `json:"count"`
	Next    string   `json:"next"`
	Prev    string   `json:"previous"`
	Results []result `json:"results"`
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	generateUsage()
	return nil
}

func generateUsage() {
	for k, _ := range commandMap {
		fmt.Printf("%s: %s\n", commandMap[k].name, commandMap[k].description)
	}
}

var con = config{URL: "https://pokeapi.co/api/v2/location-area"}

func getMap(c *config) error {

	res, err := http.Get(con.URL)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err := fmt.Sprintf("Status code:%v", res.StatusCode)
		return errors.New(err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	con = config{}

	if err := json.Unmarshal(data, &con); err != nil {
		return err
	}

	for i := 0; i < 20; i++ {
		fmt.Printf("%v\n", con.Results[i].Name)
	}

	con.URL = con.Next

	return nil
}

func getPrevMap(c *config) error {

	if con.Prev == "null" {
		fmt.Println("You're on the first page")
		return nil
	}

	con.URL = con.Prev

	getMap(&con)

	return nil
}
