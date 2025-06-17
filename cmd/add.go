package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ErebusAJ/doro/todo"
	"github.com/google/uuid"
)

type AddCommand struct {}

// name of the command
func (g *AddCommand) Name() string {
	return "add"
}

// description of the command 
func(g *AddCommand) Description() string {
	return "add creates a new to do task"
}

// run command
func(g *AddCommand) Run(args []string) error {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	tempTask := fs.String("task", "", "Your todo task.")

	// priority for the to do task
	priority := fs.Int("p", 2, "Set priority levels: 1/2/3 highest to lowest")

	fs.Parse(args)
	if *tempTask == "" {
		return fmt.Errorf("task description cannot be empty")
	}


	task := todo.TaskItem{
		ID: uuid.New().String(),
		Text: *tempTask,
		Priority: int32(*priority),
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
func saveTask(filename string, task todo.TaskItem) error {
	
	var savedTask []todo.TaskItem
	jsonb, err := os.ReadFile("./.tasks.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonb, &savedTask)
	if err != nil {
		return err
	}

	savedTask = append(savedTask, task)

	jsonItem, err := json.Marshal(savedTask)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonItem, 0644)
	if err != nil {
		return err
	}

	return nil
}
