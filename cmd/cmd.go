package cmd

import (
	"github.com/Shobhit-Nagpal/pokedexgo/internal/pokecache"
	"time"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

type MapConfig struct {
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

func GetCommands() map[string]cliCommand {
	cache := pokecache.NewCache(10 * time.Second)
	config := &MapConfig{
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
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}
