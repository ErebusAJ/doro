package todo

import "time"

// struct for each todo task item
type TaskItem struct {
	ID	 	  string	`json:"id"`
	Text 	  string	`json:"text"`
	Priority  int		`json:"priority"`
	Completed bool		`json:"completed"`
	Date	  time.Time	`json:"date"`
}