package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ErebusAJ/doro/todo"
)

type DeleteCommand struct {}

func(c *DeleteCommand) Name() string {
	return "delete"
}

func(c *DeleteCommand) Description() string {
	return "deletes a task from the to do list"
}

func(c *DeleteCommand) Run(args []string) error {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	taskID := fs.String("i", "", "task id to delete")

	err := fs.Parse(args)
	if err != nil {
		return err
	} 

	if *taskID == "" {
		return fmt.Errorf("error task id cannot be empty\n Use: \n doro delete -i <id>")
	}

	var tasks []todo.TaskItem
	temp, err := os.ReadFile(todo.TaskFilePath())
	if err != nil {
		return err
	}
	err = json.Unmarshal(temp, &tasks)
	if err != nil {
		return err
	}

	ind := -1 
	for i, t := range tasks {
		if(t.ID == *taskID) {
			ind = i
		}
	}
	if ind == -1 {
		fmt.Println("no task found with associated id")
		return nil
	} 

	updatedTask := tasks[:ind]
	updatedTask = append(updatedTask, tasks[ind + 1:]...)
	
	jsonb, err := json.Marshal(updatedTask)
	if err != nil {
		return err
	}

	err = os.WriteFile(todo.TaskFilePath(), jsonb, 0644)
	if err != nil {
		return err
	}

	fmt.Println("deleted task successfully")

	return nil
}