package cmd

import (
	"github.com/Shobhit-Nagpal/pokedexgo/internal/pokecache"
	"time"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(string) error
}

type Config struct {
	Next     string
	Previous string
	Count    int
	Cache    pokecache.Cache
}

type MapResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
	} `json:"encounter_method_rates"`
	Location struct {
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
	} `json:"pokemon_encounters"`
}

func GetCommands() map[string]cliCommand {
	cache := pokecache.NewCache(10 * time.Second)
	config := &Config{
		Cache: cache,
	}
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays names of next 20 location areas in Pokemon world",
			Callback:    createMapCommand(config),
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays names of previous 20 location areas in Pokemon world",
			Callback:    createMapbackCommand(config),
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a location area from map",
			Callback:    createExploreCommand(config),
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}
