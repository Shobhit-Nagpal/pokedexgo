package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

      parts := strings.Fields(text);
      inputCommand := parts[0]
      inputArgument := ""
      if len(parts) > 1 {
        inputArgument = parts[1]
      }
			if command, ok := commands[inputCommand]; ok {
        if command.Name == "explore" {
          if inputArgument == "" {
            fmt.Println("ERROR: Enter location area to explore")
            continue
          }
          command.Callback(inputArgument)
        } else {
          command.Callback("")
        }
			} else {
				fmt.Println("Command not recognized")
			}
		}
	}
}
