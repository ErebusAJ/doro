package cmd

import (
	"flag"
	"fmt"
)


func GreetCmd(args []string) {
	name := flag.String("name", "World", "Your name")
	flag.CommandLine.Parse(args)
	fmt.Printf("Hello, %s \n", *name)
}

func InfoCmd(args []string) {
	age := flag.Int("age", 0, "Your age")
	flag.CommandLine.Parse(args)
	
	if *age > 0 {
		fmt.Printf("You are %d years old \n", *age)
	} else {
		fmt.Printf("Age not provided! \n")
	}
}