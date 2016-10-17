package command

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

const (
	OPERATION_ID_COMMAND_EXEC = "command.exec"
)

/**
 * Execute a command using the command handler
 */

// Base class for command exec Operation
type BaseCommandExecOperation struct{}

// Id the operation
func (exec *BaseCommandExecOperation) Id() string {
	return "command.exec"
}

// Label the operation
func (exec *BaseCommandExecOperation) Label() string {
	return "Command Exec"
}

// Description for the operation
func (exec *BaseCommandExecOperation) Description() string {
	return "Execute a specified command.  This is an abstract command executor, but commands should probably add their own operations (@TODO)."
}

// Is this an internal API operation
func (exec *BaseCommandExecOperation) Internal() bool {
	return false
}
func (exec *BaseCommandExecOperation) Properties() *operation.Properties {
	return &operation.Properties{}
}
