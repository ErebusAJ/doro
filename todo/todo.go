package todo

// struct for each todo task item
type TaskItem struct {
	ID	 	  string	`json:"id"`
	Text 	  string	`json:"text"`
	Completed bool		`json:"completed"`
}