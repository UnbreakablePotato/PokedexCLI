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

// 60 seconds in nano seconds
const duration = 60000000000

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
	callbackF   func(c *config, name string) error
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

type exploreObj struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int `json:"chance"`
				ConditionValues []struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"condition_values"`
				MaxLevel int `json:"max_level"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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

var exploreCache = pokecache.NewCache(duration)

var expObj = exploreObj{}

func exploreArea(c *config, name string) error {
	fmt.Println("debug: 1")

	fullUrl := "https://pokeapi.co/api/v2/location-area" + "/" + name

	url, condition := exploreCache.Get(fullUrl)

	if !condition {
		res, err := http.Get(fullUrl)
		if err != nil {
			fmt.Printf("Request failed: %v\n", err)
			return err
		}
		fmt.Printf("debug: %s\n", fullUrl)

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Println("Area does not exist")
			err := fmt.Sprintf("Status code: %v\n", res.StatusCode)
			return errors.New(err)
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return err
		}

		if err := json.Unmarshal(data, &expObj); err != nil {
			fmt.Printf("Error: %v\n", err)
			return err
		}

		for i := range expObj.PokemonEncounters {
			fmt.Printf("%v\n", expObj.PokemonEncounters[i].Pokemon.Name)
		}

		exploreCache.Add(fullUrl, data)

		return nil
	}
	fmt.Println("Key found")
	if err := json.Unmarshal(url, &expObj); err != nil {
		return err
	}

	for i := range expObj.PokemonEncounters {
		fmt.Printf("%v\n", expObj.PokemonEncounters[i].Pokemon.Name)
	}

	return nil
}

func commandDelete(c *config, pokemon string) error {
	removePokemon(pokemon, pokemonParty)
	return nil
}
