package command

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * A command operation to retrieve a command object
 */

const (
	OPERATION_ID_COMMAND_GET = "command.get"
)

/**
 * Execute a command using the command handler
 */

// Base class for command get Operation
type BaseCommandGetOperation struct{}

// Id the operation
func (get *BaseCommandGetOperation) Id() string {
	return OPERATION_ID_COMMAND_GET
}

// Label the operation
func (get *BaseCommandGetOperation) Label() string {
	return "Command Get"
}

// Description for the operation
func (get *BaseCommandGetOperation) Description() string {
	return "Retrieve a specified command.  Retrieve a command object."
}

// Is this an internal API operation
func (get *BaseCommandGetOperation) Internal() bool {
	return false
}
func (get *BaseCommandGetOperation) Properties() *operation.Properties {
	return &operation.Properties{}
}
