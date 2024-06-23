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
	Pokedex  map[string]PokemonResponse
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

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonResponse struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience         int    `json:"base_experience"`
	Height                 int    `json:"height"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Name                   string `json:"name"`
	Order                  int    `json:"order"`
	Stats                  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
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
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetCommands() map[string]cliCommand {
	cache := pokecache.NewCache(10 * time.Second)
	config := &Config{
		Cache:   cache,
		Pokedex: map[string]PokemonResponse{},
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
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon",
			Callback:    createCatchCommand(config),
		},
		"inspect": {
			Name:        "inspect",
			Description: "Get information about the pokemon you caught",
			Callback:    createInspectCommand(config),
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "See what all pokemons you've captured",
			Callback:    createPokedexCommand(config),
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}
