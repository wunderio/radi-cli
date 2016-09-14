package command

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * A base Command definition, which defines the command
 * container configuration, but may receive overrides for
 * flags, input/error/output
 *
 * It turns out that a Command has a very similar need to
 * operations, and so it makes sense to write a command
 * interface that can be used as an operation, and can
 * give an operation
 */

// Command definition
type Command interface {
	// Attach a set of configurations to the command
	SetConfigurations(configurations *operation.Configurations)
	// Convert the command into an operation
	getOperation() operation.Operation

	/**
	 * These are taken directly from operation.Operation
	 */

	// Run a validation check on the Operation
	Validate() bool

	// Is this operation meant to be used only inside the API
	Internal() bool

	// What settings does the Operation provide to an implemenentor
	Configurations() *operation.Configurations

	// Return the string machinename/id of the Operation
	Id() string
	// Return a user readable string label for the Operation
	Label() string
	// return a multiline string description for the Operation
	Description() string

	// Execute the Operation
	Exec() operation.Result
}
