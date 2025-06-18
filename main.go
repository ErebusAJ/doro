package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/ErebusAJ/doro/cmd"
	"github.com/common-nighthawk/go-figure"
)

func printHelp(commands []cmd.Command) {
	fmt.Println("Usage: doro [command] [flags]")
	fmt.Println("\nAvailable commands: ")

	w := tabwriter.NewWriter(os.Stdout, 4, 2, 4, ' ', 0)
	for _, c := range commands {
		fmt.Fprintf(w, "\t%v\t- %v\n", c.Name(), c.Description())
	}
	w.Flush()
	fmt.Println("\nUse 'doro [command] -h' for more information about the comand.")

}

// root cmd display 
func rootCommand(commands []cmd.Command) {
	figure.NewColorFigure("Doro", "isometric1", "green", true).Print()
	figure.NewColorFigure("A CLI todo app", "ogre", "cyan", true).Print()
	printHelp(commands)
}

func main() {
	// command list
	commands := []cmd.Command{
		&cmd.AddCommand{},
		&cmd.ShowCommand{},
		&cmd.UpdateCommand{},
		&cmd.CompleteCommand{},
		&cmd.DeleteCommand{},
	}
	
	if len(os.Args) < 2 {
		rootCommand(commands)
		return
	}

	input := os.Args[1]
	for _, c := range commands {
		if input == c.Name() {
			err := c.Run(os.Args[2:])
			if err != nil {
				fmt.Printf("Error: %v \n", err)
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