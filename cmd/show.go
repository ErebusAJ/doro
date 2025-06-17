package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ErebusAJ/doro/todo"
)

type ShowCommand struct{}

func(c *ShowCommand) Name() string {
	return "show"
}


func(c *ShowCommand) Description() string {
	return "lists users to do tasks"
}


func(c *ShowCommand) Run(args []string) error {
	fs := flag.NewFlagSet("show", flag.ExitOnError)

	// show based on completion status
	status := fs.Bool("status", false, "show tasks based on completion status")

	fs.Parse(args)

	var tasks []todo.TaskItem

	jsonb, err := os.ReadFile("./.tasks.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonb, &tasks)
	if err != nil {
		return err
	}

	fmt.Println("	UUID\t\t\t\t\tTask\t\t\tPriority	Completed")
	for _, i := range tasks {
		if(i.Completed != *status){ 
			continue
		}
		fmt.Printf("	%s	%s	\t%d	\t%v \n", i.ID, i.Text,i.Priority, i.Completed)
	}

	return nil
}