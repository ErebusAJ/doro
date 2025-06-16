package main

import (
	"flag"
	"fmt"
	"os"
)

func greetCmd(args []string) {
	name := flag.String("name", "World", "Your name")
	flag.CommandLine.Parse(args)
	fmt.Printf("Hello, %s \n", *name)
}

func infoCmd(args []string) {
	age := flag.Int("age", 0, "Your age")
	flag.CommandLine.Parse(args)
	
	if *age > 0 {
		fmt.Printf("You are %d years old \n", *age)
	} else {
		fmt.Printf("Age not provided! \n")
	}
}

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
		greetCmd(os.Args[2:])
	
	case "info":
		infoCmd(os.Args[2:])
	
	case "-h", "--help":
		printHelp()

	default:
		fmt.Println("Unknown command: ", os.Args[1])
	}

}