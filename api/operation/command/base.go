package command

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Base Operation classes for command implementation
 */

// A Base command operation, that provides a command connecter
type BaseCommandConnectorOperation struct {
	connector CommandConnector
}

//
func (base *BaseCommandConnectorOperation) SetConnector(connector CommandConnector) {
	base.connector = connector
}

//
func (base *BaseCommandConnectorOperation) Connector() CommandConnector {
	return base.connector
}

// A Base command operation that returns a list of keys
type BaseCommandKeysOperation struct {
	configurations *operation.Configurations
}

// Return a static keys list configuration
func (base *BaseCommandKeysOperation) Configurations() *operation.Configurations {
	if base.configurations == nil {
		base.configurations = &operation.Configurations{}

		base.configurations.Add(operation.Configuration(&CommandKeysConfiguration{}))
	}
	return base.configurations
}

// A base class for a Command config that connects to an io.Writer
type BaseCommandWriterConfiguration struct {
	operation.WriterConfiguration
}

// A base class for a Command config that connects to an io.Writer
type BaseCommandReaderConfiguration struct {
	operation.ReaderConfiguration
}

// A base command operation that provides a a key, flags strings, and various input and output configurations
type BaseCommandKeyFlagsInputOutputOperation struct {
	configurations *operation.Configurations
}

// get static configurations
func (base *BaseCommandKeyFlagsInputOutputOperation) Configurations() *operation.Configurations {
	if base.configurations == nil {
		base.configurations = &operation.Configurations{}

		base.configurations.Add(operation.Configuration(&CommandKeyConfiguration{}))
		base.configurations.Add(operation.Configuration(&CommandFlagsConfiguration{}))
		base.configurations.Add(operation.Configuration(&CommandOutputConfiguration{}))
		base.configurations.Add(operation.Configuration(&CommandErrorConfiguration{}))
		base.configurations.Add(operation.Configuration(&CommandInputConfiguration{}))
		base.configurations.Add(operation.Configuration(&CommandContextConfiguration{}))
	}
	return base.configurations
}
