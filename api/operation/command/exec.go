package command

import (
	"errors"

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
func (exec *BaseCommandExecOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}

//
type CommandConnectorExecOperation struct {
	BaseCommandExecOperation
	BaseCommandConnectorOperation
	BaseCommandKeyFlagsInputOutputOperation
}

//
func (exec CommandConnectorExecOperation) Validate() bool {
	return true
}

//
func (exec CommandConnectorExecOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	confs := exec.BaseCommandKeyFlagsInputOutputOperation.Configurations()

	keyConf, _ := confs.Get(OPERATION_CONFIGURATION_COMMAND_KEY)

	if key, ok := keyConf.Get().(string); ok {
		if keyedCommand, ok := exec.Connector().Get(key); ok {
			keyedCommand.SetConfigurations(confs)
			success, errs := keyedCommand.Exec().Success()
			result.Set(success, errs)
		} else {
			result.Set(false, []error{errors.New("Command connector did not find the command you were looking for")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not get a command from the command connector")})
	}

	return operation.Result(&result)
}
