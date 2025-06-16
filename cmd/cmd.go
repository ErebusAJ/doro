package cmd

// interface for commands
type Command interface {
	Name() string
	Description() string
	Run(args []string) error
}