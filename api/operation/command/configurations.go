package command

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Typical command package operation configurations.
 */

const (
	// Key string for a single operation
	OPERATION_CONFIGURATION_COMMAND_KEY = "command.key"
	// List of keys
	OPERATION_CONFIGURATION_COMMAND_KEYS = "command.keys"

	// list of string flags passed to the command container
	OPERATION_CONFIGURATION_COMMAND_FLAGS = "command.flags"

	// Input/Output objects
	OPERATION_CONFIGURATION_COMMAND_OUTPUT = "command.output"
	OPERATION_CONFIGURATION_COMMAND_ERR    = "command.err"
	OPERATION_CONFIGURATION_COMMAND_INPUT  = "command.input"

	// Use a context when running, to allow remote control of execution
	OPERATION_CONFIGURATION_COMMAND_CONTEXT = "command.context"
)

// Command for a single command key
type CommandKeyConfiguration struct {
	operation.StringConfiguration
}

// Id for the configuration
func (confKey *CommandKeyConfiguration) Id() string {
	return OPERATION_CONFIGURATION_COMMAND_KEY
}

// Label for the configuration
func (confKey *CommandKeyConfiguration) Label() string {
	return "Command key."
}

// Description for the configuration
func (confKey *CommandKeyConfiguration) Description() string {
	return "Command key."
}

// Command for an ordered list of command keys
type CommandKeysConfiguration struct {
	operation.StringSliceConfiguration
}

// Id for the configuration
func (keyValue *CommandKeysConfiguration) Id() string {
	return OPERATION_CONFIGURATION_COMMAND_KEYS
}

// Label for the configuration
func (keyValue *CommandKeysConfiguration) Label() string {
	return "Command key list."
}

// Description for the configuration
func (keyValue *CommandKeysConfiguration) Description() string {
	return "Command key list."
}

// Command for an ordered list of command keys
type CommandFlagsConfiguration struct {
	operation.StringSliceConfiguration
}

// Id for the configuration
func (keyValue *CommandFlagsConfiguration) Id() string {
	return OPERATION_CONFIGURATION_COMMAND_FLAGS
}

// Label for the configuration
func (keyValue *CommandFlagsConfiguration) Label() string {
	return "Command flags list."
}

// Description for the configuration
func (keyValue *CommandFlagsConfiguration) Description() string {
	return "An ordered list of string flags to send to a command."
}

// A command configuration for command output
type CommandOutputConfiguration struct {
	operation.WriterConfiguration
}

// Id for the configuration
func (keyValue *CommandOutputConfiguration) Id() string {
	return OPERATION_CONFIGURATION_COMMAND_OUTPUT
}

// Label for the configuration
func (keyValue *CommandOutputConfiguration) Label() string {
	return "Command output io.Writer."
}

// Description for the configuration
func (keyValue *CommandOutputConfiguration) Description() string {
	return "An io.Writer, which will receive the command execution output.  Any io.writer can be used, the default here will be os.Stdout."
}

// A command configuration for command error output
type CommandErrorConfiguration struct {
	BaseCommandWriterConfiguration
}

// Id for the configuration
func (keyValue *CommandErrorConfiguration) Id() string {
	return OPERATION_CONFIGURATION_COMMAND_ERR
}

// Label for the configuration
func (keyValue *CommandErrorConfiguration) Label() string {
	return "Command error io.Writer."
}

// Description for the configuration
func (keyValue *CommandErrorConfiguration) Description() string {
	return "An io.Writer, which will receive the command execution error output.  Any io.writer can be used, the default here will be os.Stdout."
}

// A command configuration for command execution input
type CommandInputConfiguration struct {
	BaseCommandReaderConfiguration
}

// Id for the configuration
func (keyValue *CommandInputConfiguration) Id() string {
	return OPERATION_CONFIGURATION_COMMAND_INPUT
}

// Label for the configuration
func (keyValue *CommandInputConfiguration) Label() string {
	return "Command input io.Reader."
}

// Description for the configuration
func (keyValue *CommandInputConfiguration) Description() string {
	return "An io.Reader, which will provide command execution input.  Any io.reader can be used, the default here will be os.Stdin"
}

// A command configuration for command execution net context
type CommandContextConfiguration struct {
	operation.ContextConfiguration
}

// Id for the configuration
func (contextConf *CommandContextConfiguration) Id() string {
	return OPERATION_CONFIGURATION_COMMAND_CONTEXT
}

// Label for the configuration
func (contextConf *CommandContextConfiguration) Label() string {
	return "Command context limiter"
}

// Description for the configuration
func (contextConf *CommandContextConfiguration) Description() string {
	return "A golang.org/x/net/context for controling execution."
}
