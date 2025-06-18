package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ErebusAJ/doro/todo"
)

type UpdateCommand struct {}

func(c *UpdateCommand) Name() string {
	return "update"
}

func(c * UpdateCommand) Description() string {
	return "update any task description or priority"
}

func(c *UpdateCommand) Run(args []string) error {
	fs := flag.NewFlagSet("update", flag.ExitOnError)
	id := fs.String("i", "", "id of task to update")
	text := fs.String("task", "", "todo task updated description")
	priority := fs.Int("p", 0, "todo task updated priority")

	fs.Parse(args)

	if *id == "" {
		return fmt.Errorf("i flag for id cannot be empty")
	}

	if *priority != 1 && *priority != 2 {
		return fmt.Errorf("priority can take values 1/2/3")
	}

	var tasks []todo.TaskItem
	jsonb, err := os.ReadFile(todo.TaskFilePath())
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonb, &tasks)
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if(t.ID == *id) {
			if(*text != "") {
				tasks[i].Text = *text
			}
			if(*priority != 0) {
				tasks[i].Priority = *priority
			}
		}
	}

	jsonb, err = json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile(todo.TaskFilePath(), jsonb, 0644)
	if err != nil {
		return err
	}

	
	return nil
}