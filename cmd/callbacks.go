package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const URL = "https://pokeapi.co/api/v2/location/"

func commandHelp() error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Displays names of next 20 location areas in Pokemon world
mapb: Displays names of previous 20 location areas in Pokemon world`)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func createMapCommand(config *MapConfig) func() error {
	return func() error {
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

func createMapbackCommand(config *MapConfig) func() error {
	return func() error {
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
