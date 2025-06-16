package main

import (
	"fmt"
	"os"

	"github.com/ErebusAJ/go-cli/cmd"
)

func printHelp() {
	fmt.Println("Usage: go-cli [command] [flags]")
	fmt.Println("\nAvailablle commands: ")
	fmt.Println("	greet	- Greet someone")
	fmt.Println("	info	- Print age")
	fmt.Println("\nUse 'go-cli [command] -h' for more information about the comand.")

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-cli [command] [flags]")
		return
	}

	switch os.Args[1] {
	case "greet":
		cmd.GreetCmd(os.Args[2:])
	
	case "info":
		cmd.InfoCmd(os.Args[2:])
	
	case "-h", "--help":
		printHelp()

	default:
		fmt.Println("Unknown command: ", os.Args[1])
	}

}