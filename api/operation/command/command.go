package command

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * A set of commands
 */

type Commands struct {
	commands map[string]Command
	order    []string
}

// Safe lazy constructor
func (commands *Commands) safe() {
	if &commands.commands == nil {
		commands.commands = map[string]Command{}
		commands.order = []string{}
	}
}

// Get a command
func (commands *Commands) Get(key string) (Command, bool) {
	commands.safe()
	comm, found := commands.commands[key]
	return comm, found
}

// Add a command
func (commands *Commands) Set(key string, comm Command) error {
	commands.safe()
	if _, exists := commands.commands[key]; !exists {
		commands.order = append(commands.order, key)
	}
	commands.commands[key] = comm
	return nil
}

// Order of commands
func (commands *Commands) Order() []string {
	commands.safe()
	return commands.order
}

/**
 * A base Command definition, which defines the command
 * container property, but may receive overrides for
 * flags, input/error/output
 *
 * It turns out that a Command has a very similar need to
 * operations, and so it makes sense to write a command
 * interface that can be used as an operation, and can
 * give an operation
 */

// Command definition
type Command interface {
	// Return the string machinename/id of the Operation
	Id() string
	// Return a user readable string label for the Operation
	Label() string
	// return a multiline string description for the Operation
	Description() string

	// Run a validation check on the Operation
	Validate() bool

	// Is this operation meant to be used only inside the API
	Internal() bool

	// What settings does the Operation provide to an implemenentor
	Properties() *operation.Properties

	// Execute the Operation
	Exec() operation.Result
}
