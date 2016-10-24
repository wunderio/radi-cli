package command

import (
	log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Typical command package operation Properies.
 */

const (
	// Key string for a single operation
	OPERATION_PROPERTY_COMMAND_KEY = "command.key"
	// Command object for a single operation
	OPERATION_PROPERTY_COMMAND_COMMAND = "command.command"
	// List of keys
	OPERATION_PROPERTY_COMMAND_KEYS = "command.keys"

	// list of string flags passed to the command container
	OPERATION_PROPERTY_COMMAND_FLAGS = "command.flags"

	// Input/Output objects
	OPERATION_PROPERTY_COMMAND_OUTPUT = "command.output"
	OPERATION_PROPERTY_COMMAND_ERR    = "command.err"
	OPERATION_PROPERTY_COMMAND_INPUT  = "command.input"

	// Use a context when running, to allow remote control of execution
	OPERATION_PROPERTY_COMMAND_CONTEXT = "command.context"
)

// Command for a single command key
type CommandKeyProperty struct {
	operation.StringProperty
}

// Id for the Property
func (confKey *CommandKeyProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_KEY
}

// Label for the Property
func (confKey *CommandKeyProperty) Label() string {
	return "Command key."
}

// Description for the Property
func (confKey *CommandKeyProperty) Description() string {
	return "Command key."
}

// Is the Property internal only
func (confKey *CommandKeyProperty) Internal() bool {
	return false
}

// Command for a single command object
type CommandCommandProperty struct {
	command Command
}

// Id for the Property
func (com *CommandCommandProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_COMMAND
}

// Label for the Property
func (com *CommandCommandProperty) Label() string {
	return "Command object."
}

// Description for the Property
func (com *CommandCommandProperty) Description() string {
	return "Command object."
}

// Is the Property internal only
func (com *CommandCommandProperty) Internal() bool {
	return true
}

// Give an idea of what type of value the property consumes
func (com *CommandCommandProperty) Type() string {
	return "operation/command.Command"
}

// Get the Command value
func (com *CommandCommandProperty) Get() interface{} {
	return interface{}(com.command)
}

// Set the Command value
func (com *CommandCommandProperty) Set(value interface{}) bool {
	if converted, ok := value.(Command); ok {
		com.command = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected Command")
		return false
	}
}

// Command for an ordered list of command keys
type CommandKeysProperty struct {
	operation.StringSliceProperty
}

// Id for the Property
func (keyValue *CommandKeysProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_KEYS
}

// Label for the Property
func (keyValue *CommandKeysProperty) Label() string {
	return "Command key list."
}

// Description for the Property
func (keyValue *CommandKeysProperty) Description() string {
	return "Command key list."
}

// Is the Property internal only
func (keyValue *CommandKeysProperty) Internal() bool {
	return false
}

// Command for an ordered list of command keys
type CommandFlagsProperty struct {
	operation.StringSliceProperty
}

// Id for the Property
func (keyValue *CommandFlagsProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_FLAGS
}

// Label for the Property
func (keyValue *CommandFlagsProperty) Label() string {
	return "Command flags list."
}

// Description for the Property
func (keyValue *CommandFlagsProperty) Description() string {
	return "An ordered list of string flags to send to a command."
}

// Is the Property internal only
func (keyValue *CommandFlagsProperty) Internal() bool {
	return false
}

// A command Property for command output
type CommandOutputProperty struct {
	operation.WriterProperty
}

// Id for the Property
func (keyValue *CommandOutputProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_OUTPUT
}

// Label for the Property
func (keyValue *CommandOutputProperty) Label() string {
	return "Command output io.Writer."
}

// Description for the Property
func (keyValue *CommandOutputProperty) Description() string {
	return "An io.Writer, which will receive the command execution output.  Any io.writer can be used, the default here will be os.Stdout."
}

// Is the Property internal only
func (keyValue *CommandOutputProperty) Internal() bool {
	return false
}

// A command Property for command error output
type CommandErrorProperty struct {
	operation.WriterProperty
}

// Id for the Property
func (keyValue *CommandErrorProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_ERR
}

// Label for the Property
func (keyValue *CommandErrorProperty) Label() string {
	return "Command error io.Writer."
}

// Description for the Property
func (keyValue *CommandErrorProperty) Description() string {
	return "An io.Writer, which will receive the command execution error output.  Any io.writer can be used, the default here will be os.Stdout."
}

// Is the Property internal only
func (keyValue *CommandErrorProperty) Internal() bool {
	return false
}

// A command Property for command execution input
type CommandInputProperty struct {
	operation.ReaderProperty
}

// Id for the Property
func (keyValue *CommandInputProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_INPUT
}

// Label for the Property
func (keyValue *CommandInputProperty) Label() string {
	return "Command input io.Reader."
}

// Description for the Property
func (keyValue *CommandInputProperty) Description() string {
	return "An io.Reader, which will provide command execution input.  Any io.reader can be used, the default here will be os.Stdin"
}

// Is the Property internal only
func (keyValue *CommandInputProperty) Internal() bool {
	return false
}

// A command Property for command execution net context
type CommandContextProperty struct {
	operation.ContextProperty
}

// Id for the Property
func (contextConf *CommandContextProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_CONTEXT
}

// Label for the Property
func (contextConf *CommandContextProperty) Label() string {
	return "Command context limiter"
}

// Description for the Property
func (contextConf *CommandContextProperty) Description() string {
	return "A golang.org/x/net/context for controling execution."
}

// Is the Property internal only
func (contextConf *CommandContextProperty) Internal() bool {
	return false
}
