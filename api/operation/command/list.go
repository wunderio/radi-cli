package command

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * An operation for listing commands that are available
 * in the app
 */

const (
	OPERATION_ID_COMMAND_LIST = "command.list"
)

// Base class for command list Operation
type BaseCommandListOperation struct{}

// Id the operation
func (list *BaseCommandListOperation) Id() string {
	return "command.list"
}

// Label the operation
func (list *BaseCommandListOperation) Label() string {
	return "Command List"
}

// Description for the operation
func (list *BaseCommandListOperation) Description() string {
	return "List all available commands."
}

// Is this an internal API operation
func (list *BaseCommandListOperation) Internal() bool {
	return true
}
func (list *BaseCommandListOperation) Properties() *operation.Properties {
	return &operation.Properties{}
}
