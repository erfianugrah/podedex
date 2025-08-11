package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(urls *config, args []string) error
}

type config struct {
	NextURL     *string
	PreviousURL *string
	Cache       *pokecache.Cache
}

func main() {
	cache := pokecache.NewCache(5 * time.Second)
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
		"explore": {
			name:        "Explore a location",
			description: "Location details",
			callback:    commandExplore,
		},
	}

	apiConfig := &config{
		NextURL:     nil,
		PreviousURL: nil,
		Cache:       cache,
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		output := cleanInput(input)

		if len(output) > 0 {
			commandName := output[0]
			args := output[1:]
			if command, exists := commands[commandName]; exists {
				err := command.callback(apiConfig, args)
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

func commandExit(urls *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(urls *config, args []string) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:
  help: Displays a help message
  exit: Exit the Pokedex
`)
	return nil
}

func commandMap(urls *config, args []string) error {
	var urlToFetch string
	if urls.NextURL == nil {
		urlToFetch = pokeapi.BaseURL
	} else {
		urlToFetch = *urls.NextURL
	}
	cachedData, found := urls.Cache.Get(urlToFetch)
	if found {
		fmt.Println("Using cached data!") // Optional: show when cache is used
		var res pokeapi.Location
		if err := json.Unmarshal(cachedData, &res); err != nil {
			return err
		}
		urls.NextURL = res.Next
		urls.PreviousURL = res.Previous
		for i := 0; i < len(res.Results); i++ {
			fmt.Println(res.Results[i].Name)
		}
		return nil

	}
	fmt.Println("Fetching from API...") // Optional: show when making API call
	res, err := pokeapi.FetchLocation(urlToFetch)
	if err != nil {
		return err
	}

	// Add to cache (marshal the result)
	data, err := json.Marshal(res)
	if err == nil {
		urls.Cache.Add(urlToFetch, data)
	}

	urls.NextURL = res.Next
	urls.PreviousURL = res.Previous

	for i := 0; i < len(res.Results); i++ {
		fmt.Println(res.Results[i].Name)
	}
	return nil
}

func commandMapBack(urls *config, args []string) error {
	var urlToFetch string
	if urls.PreviousURL == nil {
		fmt.Println("You're on the first page")
		return nil
	} else {
		urlToFetch = *urls.PreviousURL

		cachedData, found := urls.Cache.Get(urlToFetch)

		if found {
			fmt.Println("Using cached data!") // Optional: show when cache is used
			var res pokeapi.Location
			if err := json.Unmarshal(cachedData, &res); err != nil {
				return err
			}
			urls.NextURL = res.Next
			urls.PreviousURL = res.Previous
			for i := 0; i < len(res.Results); i++ {
				fmt.Println(res.Results[i].Name)
			}
			return nil

		}
		fmt.Println("Fetching from API...") // Optional: show when making API call
		res, err := pokeapi.FetchLocation(urlToFetch)
		if err != nil {
			return err
		}

		// Add to cache (marshal the result)
		data, err := json.Marshal(res)
		if err == nil {
			urls.Cache.Add(urlToFetch, data)
		}
		urls.NextURL = res.Next
		urls.PreviousURL = res.Previous

		for i := 0; i < len(res.Results); i++ {
			fmt.Println(res.Results[i].Name)
		}
	}
	return nil
}

func commandExplore(urls *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please input a location")
		return nil
	}
	locationName := args[0]
	url := pokeapi.BaseURL + locationName

	fmt.Printf("Exploring %s...\n", locationName)

	cachedData, found := urls.Cache.Get(url)

	if found {
		fmt.Println("Using cached data!") // Optional: show when cache is used
		var res pokeapi.LocationAreaResponse
		if err := json.Unmarshal(cachedData, &res); err != nil {
			return err
		}
		pokemonMap := make(map[string]bool)
		for i := 0; i < len(res.PokemonEncounters); i++ {
			pokemonName := res.PokemonEncounters[i].Pokemon.Name
			pokemonMap[pokemonName] = true
		}
		fmt.Println("Found Pokemon:")
		for pokemonName := range pokemonMap {
			fmt.Printf(" - %s\n", pokemonName)
		}
		return nil

	}
	fmt.Println("Fetching from API...") // Optional: show when making API call
	res, err := pokeapi.FetchLocationArea(url)
	if err != nil {
		return err
	}

	// Add to cache (marshal the result)
	data, err := json.Marshal(res)
	if err == nil {
		urls.Cache.Add(url, data)
	}

	pokemonMap := make(map[string]bool)
	for i := 0; i < len(res.PokemonEncounters); i++ {
		pokemonName := res.PokemonEncounters[i].Pokemon.Name
		pokemonMap[pokemonName] = true
	}
	fmt.Println("Found Pokemon:")
	for pokemonName := range pokemonMap {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil

}
