package cmd

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
