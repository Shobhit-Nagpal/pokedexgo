package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const URL = "https://pokeapi.co/api/v2"

func commandHelp(s string) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Displays names of next 20 location areas in Pokemon world
mapb: Displays names of previous 20 location areas in Pokemon world
inspect: Get information about pokemon you've caught
pokedex: See all pokemons you've caught
explore: Explores area given and finds pokemons present in that area`)
	return nil
}

func commandExit(s string) error {
	os.Exit(0)
	return nil
}

func createMapCommand(config *Config) func(string) error {
	return func(s string) error {
		config.Count++
		if config.Next == "" {
			data, body, err := getLocation(URL)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			config.Cache.Add("1", body)
			config.Next = data.Next
			if data.Previous == nil {
				config.Previous = ""
			} else {
				config.Previous = *data.Previous
			}

			printLocations(data)

		} else {
			val, ok := config.Cache.Get(strconv.Itoa(config.Count))
			if ok {
				data := MapResponse{}
				err := json.Unmarshal(val, &data)
				if err != nil {
					return err
				}
				config.Next = data.Next
				if data.Previous == nil {
					config.Previous = ""
				} else {
					config.Previous = *data.Previous
				}
				printLocations(data)
				return nil
			}

			data, body, err := getLocation(config.Next)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			config.Cache.Add(strconv.Itoa(config.Count), body)

			config.Next = data.Next
			if data.Previous == nil {
				config.Previous = ""
			} else {
				config.Previous = *data.Previous
			}

			printLocations(data)
		}
		return nil
	}
}

func createMapbackCommand(config *Config) func(string) error {
	return func(s string) error {
		if config.Previous == "" {
			fmt.Println("ERROR: You're on first page of results. Run map command first")
			return errors.New("You're on first page of results")
		}

		config.Count--
		val, ok := config.Cache.Get(strconv.Itoa(config.Count))
		if ok {
			data := MapResponse{}
			err := json.Unmarshal(val, &data)
			if err != nil {
				return err
			}
			printLocations(data)
			config.Next = data.Next
			if data.Previous == nil {
				config.Previous = ""
			} else {
				config.Previous = *data.Previous
			}
			return nil
		}

		data, body, err := getLocation(config.Previous)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		config.Cache.Add(strconv.Itoa(config.Count), body)
		printLocations(data)

		config.Next = data.Next
		if data.Previous == nil {
			config.Previous = ""
		} else {
			config.Previous = *data.Previous
		}
		return nil
	}
}

func createExploreCommand(config *Config) func(string) error {
	return func(area string) error {
		fmt.Printf("Exploring %s...\n", area)
		if val, ok := config.Cache.Get("area"); ok {
			data := LocationAreaResponse{}
			err := json.Unmarshal(val, &data)
			if err != nil {
				return err
			}
			printPokemons(data)
		} else {
			data, body, err := getPokemonsInArea(URL, area)
			if err != nil {
				return err
			}
			config.Cache.Add(area, body)

			printPokemons(data)
		}
		return nil
	}
}

func createCatchCommand(config *Config) func(string) error {
	return func(pokemon string) error {
    pokemon = strings.ToLower(pokemon)
    if _, exists := config.Pokedex[pokemon]; exists {
      fmt.Printf("%s has already been caught\n", pokemon)
      return nil
    }
		fmt.Printf("Throwing a pokeball at %s...\n", pokemon)

    data, _, err := getPokemonInfo(URL, pokemon)
    if err != nil {
      return err
    }
    baseExp := data.BaseExperience
    randomInt := rand.Intn(baseExp + 69)
    if randomInt > baseExp {
      fmt.Printf("%s was caught!\n", pokemon)
      fmt.Println("You may now inspect it with the inspect command")
      config.Pokedex[pokemon] = data
    } else {
      fmt.Printf("%s escaped!\n", pokemon)
    }
    return nil
  }
}

func createInspectCommand(config *Config) func(string) error {
  return func(pokemon string) error {
    pokemon = strings.ToLower(pokemon)

    if data, exists := config.Pokedex[pokemon]; exists {
      printPokemonInfo(data)
    } else {
      fmt.Println("You have not caught that pokemon")
    }
    return nil
  }
}

func createPokedexCommand(config *Config) func(string) error {
  return func(s string) error {
    fmt.Println("Your pokedex:")
    if len(config.Pokedex) == 0 {
      fmt.Println("You have not caught any pokemons")
      return nil
    }

    for _, pokemon := range config.Pokedex {
      fmt.Printf("\t - %s\n", pokemon.Name)
    }
    return nil
  }
}
