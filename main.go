package main

import (
	"fmt"
	"os"

	"github.com/ErebusAJ/doro/cmd"
)

func printHelp(commands []cmd.Command) {
	fmt.Println("Usage: doro [command] [flags]")
	fmt.Println("\nAvailablle commands: ")
	for _, c := range commands {
		fmt.Printf("	%v	- %v", c.Name(), c.Description())
	}
	fmt.Println("\nUse 'doro [command] -h' for more information about the comand.")

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: doro [command] [flags]")
		return
	}

	commands := []cmd.Command{
		&cmd.AddCommand{},
		&cmd.ShowCommand{},
	}


	input := os.Args[1]
	for _, c := range commands {
		if input == c.Name() {
			err := c.Run(os.Args[2:])
			if err != nil {
				fmt.Printf("Error: %v ", err)
				os.Exit(1)
			}
			return
		}
	}

	switch os.Args[1] {
	case "greet":
		cmd.GreetCmd(os.Args[2:])
	
	case "info":
		cmd.InfoCmd(os.Args[2:])
	
	case "-h", "--help":
		printHelp(commands)

	default:
		fmt.Println("Unknown command: ", os.Args[1])
	}

}