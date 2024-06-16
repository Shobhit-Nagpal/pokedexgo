package cmd

import (
	"errors"
	"fmt"
	"os"
)

const URL = "https://pokeapi.co/api/v2/location/"

func commandHelp() error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Displays names of next 20 location areas in Pokemon world
mapb: Displays names of previous 20 location areas in Pokemon world
`)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func createMapCommand(config *MapConfig) func() error {
	return func() error {
		if config.Next == "" {
			data, err := getLocation(URL)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			config.Next = data.Next
			if data.Previous == nil {
				config.Previous = ""
			} else {
				config.Previous = *data.Previous
			}

			for _, location := range data.Results {
				fmt.Println(location.Name)
			}

		} else {
			data, err := getLocation(config.Next)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			config.Next = data.Next
			if data.Previous == nil {
				config.Previous = ""
			} else {
				config.Previous = *data.Previous
			}

			for _, location := range data.Results {
				fmt.Println(location.Name)
			}
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

		data, err := getLocation(config.Previous)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		for _, location := range data.Results {
			fmt.Println(location.Name)
		}

		config.Next = data.Next
		if data.Previous == nil {
			config.Previous = ""
		} else {
			config.Previous = *data.Previous
		}
		return nil
	}
}
