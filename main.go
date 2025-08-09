package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(urls *config) error
}

type config struct {
	NextURL     *string
	PreviousURL *string
}

type location struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}
type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

const BASE_URL = "https://pokeapi.co/api/v2/location-area/"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := map[string]cliCommand{
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
			description: "Displays next list",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "Displays previous list",
			callback:    commandMapBack,
		},
	}

	apiConfig := &config{
		NextURL:     nil,
		PreviousURL: nil,
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		output := cleanInput(input)

		if len(output) > 0 {
			commandName := output[0]
			if command, exists := commands[commandName]; exists {
				err := command.callback(apiConfig)
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

func fetchLocation(url string) (*location, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if req.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", req.StatusCode)
	}

	data, err := io.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	var newLocation location
	if err := json.Unmarshal(data, &newLocation); err != nil {
		return nil, err
	}
	return &newLocation, nil
}

func commandExit(urls *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(urls *config) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:
  help: Displays a help message
  exit: Exit the Pokedex
`)
	return nil
}

func commandMap(urls *config) error {
	var urlToFetch string
	if urls.NextURL == nil {
		urlToFetch = BASE_URL
	} else {
		urlToFetch = *urls.NextURL
	}
	res, err := fetchLocation(urlToFetch)
	if err != nil {
		return err
	}
	urls.NextURL = res.Next
	urls.PreviousURL = res.Previous

	for i := 0; i < len(res.Results); i++ {
		fmt.Println(res.Results[i].Name)
	}
	return nil
}

func commandMapBack(urls *config) error {
	var urlToFetch string
	if urls.PreviousURL == nil {
		fmt.Println("You're on the first page")
		return nil
	} else {
		urlToFetch = *urls.PreviousURL
		res, err := fetchLocation(urlToFetch)
		if err != nil {
			return err
		}
		urls.NextURL = res.Next
		urls.PreviousURL = res.Previous

		for i := 0; i < len(res.Results); i++ {
			fmt.Println(res.Results[i].Name)
		}
	}

	return nil
}
