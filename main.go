package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/Shobhit-Nagpal/pokedexgo/cmd"
)

func main() {
	commands := cmd.GetCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); ok {
			text := scanner.Text()
			if text == "" {
				continue
			}

			if command, ok := commands[text]; ok {
        command.Callback()
			} else {
				fmt.Println("Command not recognized")
			}
		}
	}
}
