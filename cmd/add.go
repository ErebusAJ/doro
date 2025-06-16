package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ErebusAJ/go-cli/todo"
	"github.com/google/uuid"
)

type AddCommand struct {}

// name of the command
func (g *AddCommand) Name() string {
	return "add"
}

// description of the command 
func(g *AddCommand) Description() string {
	return "Add creates a new to do task."
}

// run command
func(g *AddCommand) Run(args []string) error {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	tempTask := fs.String("task", "", "Your todo task.")

	fs.Parse(args)
	if *tempTask == "" {
		return fmt.Errorf("task description cannot be empty")
	}

	task := todo.TaskItem{
		ID: uuid.New().String(),
		Text: *tempTask,
		Completed: false,
	}

	err := saveTask("./.tasks.json", task)
	if err != nil {
		return err
	}

	fmt.Println(task)

	return nil
}



// save task 
func saveTask(filename string, items todo.TaskItem) error {
	jsonItem, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonItem, 0644)
	if err != nil {
		return err
	}

	return nil
}
