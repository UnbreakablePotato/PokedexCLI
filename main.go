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
		res := cleanInput(text)

		value, ok := commandMap[res[0]]

		if ok {
			if res[0] == value.name {
				commandMap[value.name].callback(&emptyStruct)

			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
