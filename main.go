package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"net/http"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	Next string
	Previous string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:   "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:"map",
			description: "Displays a list of maps",
			callback: map,
		},
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		output := cleanInput(input)

		if len(output) > 0 {
			commandName := output[0]
			if command, exists := commands[commandName]; exists {
				err := command.callback()
				if err != nil {
					fmt.Print(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {
	formattedText := strings.Fields(strings.ToLower(text))
	return formattedText
}

func commandExit(urls config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(urls config) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:
  help: Displays a help message
  exit: Exit the Pokedex
`)
	return nil
}

func map(urls config) error {
	
}

func mapb(urls config) error {

}
