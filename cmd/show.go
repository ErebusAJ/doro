package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"

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

	// show based on priority
	shouldSort := fs.String("s", "n", "sort tasks according to priority - h/l")

	fs.Parse(args)

	// sort input checking
	if *shouldSort != "h" && *shouldSort != "l" && *shouldSort != "n" {
		return fmt.Errorf("s sort flag can only take h/l as input")
	} 

	var tasks []todo.TaskItem
	jsonb, err := os.ReadFile("./.tasks.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonb, &tasks)
	if err != nil {
		return err
	}

	// create a new writer for consistent output
	w := tabwriter.NewWriter(os.Stdout, 8, 2, 4,' ', 0)
	fmt.Fprintln(w, "Priority\tTask\tStatus\tDate Added\tID")

	// sort task based on priority 
	if *shouldSort == "h" {
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].Priority > tasks[j].Priority
		})
	} else if *shouldSort == "l" {
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].Priority < tasks[j].Priority
		})
	}

	for _, i := range tasks {
		if(i.Completed != *status){ 
			continue
		}
		fmt.Fprintln(w, strconv.Itoa(i.Priority)+"\t"+i.Text+"\t"+prettyStatus(i.Completed)+"\t"+i.Date.Format("02-01-2006")+"\t"+i.ID)
	}
	w.Flush()

	return nil
}


// beautify completion satus
func prettyStatus(status bool) string {
	if status {
		return "✅"
	}

	return "❌"
}