package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ErebusAJ/doro/todo"
)

type CompleteCommand struct {}

func(c *CompleteCommand) Name() string {
	return "complete"
}

func(c *CompleteCommand) Description() string {
	return "change the completion status of a task"
}

func(c *CompleteCommand) Run(args []string) error {
	fs := flag.NewFlagSet("complete", flag.ExitOnError)

	id := fs.String("i", "", "task id")
	fs.Parse(args)

	if *id == "" {
		return fmt.Errorf("error task id cannot be empty")
	}

	// read tasks from file
	var tasks []todo.TaskItem
	jsonb, err := os.ReadFile("./.tasks.json")
	if err != nil {
		return err
	}
	json.Unmarshal(jsonb, &tasks)

	// change status
	var updatedTask todo.TaskItem
	for ind, task := range tasks {
		if(task.ID == *id) {
			updatedTask = task
			tasks[ind].Completed = true
		}
	}

	// update file
	jsonb, err = json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile("./.tasks.json", jsonb, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Status updated: completed task - %v\n", updatedTask.Text)
	return nil
}