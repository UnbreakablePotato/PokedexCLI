package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/UnbreakablePotato/pokedexcli/internal/pokecache"
)

const duration = 60000000000

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

var cache = pokecache.NewCache(duration)

var con = config{URL: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"}

func getMap(c *config) error {

	keyStr := con.URL

	fmt.Printf("Current con.URL: %s\n", keyStr)

	val, b := cache.Get(keyStr)

	fmt.Println("In map function...")

	if b == true {
		fmt.Println("debug: Found key in cache...")
		if err := json.Unmarshal(val, &con); err != nil {
			return err
		}

		fmt.Println("debug: No error when unmarshalling json...")

		for i := 0; i < 20; i++ {
			fmt.Printf("%v\n", con.Results[i].Name)
		}

		con.URL = con.Next

		return nil
	} else {
		fmt.Println("debug: Not in cache, making GET request")
		res, err := http.Get(con.URL)
		fmt.Println("debug: Request successful...")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return err
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			err := fmt.Sprintf("Status code:%v", res.StatusCode)
			return errors.New(err)
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return err
		}

		cache.Add(con.URL, data)

		con = config{}

		fmt.Println("debug: ReadAll sucessful")

		if err := json.Unmarshal(data, &con); err != nil {
			fmt.Printf("Error: %v\n", err)
			return err
		}

		fmt.Println("debug: Unmarshal successful")

		for i := 0; i < 20; i++ {
			fmt.Printf("%v\n", con.Results[i].Name)
		}

		con.URL = con.Next

		return nil
	}
}

func getPrevMap(c *config) error {

	if con.Prev == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	con.URL = con.Prev

	getMap(&con)

	return nil
}
