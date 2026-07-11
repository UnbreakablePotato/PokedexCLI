package main

import (
	"bufio"
	"fmt"
	"os"
)

var commandMap map[string]cliCommand

func main() {

	emptyStruct := config{
		Next: "",
		Prev: "",
	}

	commandMap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists 20 location areas in the Pokemon world",
			callback:    getMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 location areas in the Pokemon world",
			callback:    getPrevMap,
		},
		"explore": {
			name:        "explore",
			description: "Lists all the Pokemon in a location area",
			callbackF:   exploreArea,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon",
			callbackF:   catch,
		},
		"inspect": {
			name:        "inspect",
			description: "Shows",
			callbackF:   inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all caught pokemons",
			callback:    printPokedex,
		},
		"delete": {
			name:        "delete",
			description: "Deletes a specific pokemon from your party",
			callbackF:   commandDelete,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input := scanner.Scan()

		if !input {
			fmt.Println("You managed to cook my repl")
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("You managed to cook my repl")
			fmt.Fprintln(os.Stderr, err)
		}

		text := scanner.Text()
		res := []string{}
		res = append(res, cleanInput(text)...)

		if len(res) < 2 {
			value, ok := commandMap[res[0]]

			if ok {
				if res[0] == value.name {
					commandMap[value.name].callback(&emptyStruct)

				}
			} else {
				fmt.Println("Unknown command")
			}
		} else if len(res) == 2 {
			command, ok := commandMap[res[0]]
			fmt.Println("debug: Entered command with 2 inputs")
			if !ok {
				fmt.Println("Unknown command")
			}
			commandMap[command.name].callbackF(&emptyStruct, res[1])
			fmt.Println("debug: callbackF called")
		}

	}
}
