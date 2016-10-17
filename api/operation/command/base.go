package command

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Base Operation classes for command implementation
 */

// A Base command operation, that provides a command connecter
// type BaseCommandConnectorOperation struct {
// 	connector CommandConnector
// }

// //
// func (base *BaseCommandConnectorOperation) SetConnector(connector CommandConnector) {
// 	base.connector = connector
// }

// //
// func (base *BaseCommandConnectorOperation) Connector() CommandConnector {
// 	return base.connector
// }

// A Base command operation that returns a list of keys
type BaseCommandKeysOperation struct {
	properties *operation.Properties
}

// Return a static keys list Property
func (base *BaseCommandKeysOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&CommandKeysProperty{}))
	}
	return base.properties
}

// A base class for a Command config that connects to an io.Writer
type BaseCommandWriterProperty struct {
	operation.WriterProperty
}

// A base class for a Command config that connects to an io.Writer
type BaseCommandReaderProperty struct {
	operation.ReaderProperty
}

// A base command operation that provides a a key, flags strings, and various input and output Properties
type BaseCommandKeyFlagsInputOutputOperation struct {
	properties *operation.Properties
}

// get static Properties
func (base *BaseCommandKeyFlagsInputOutputOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&CommandKeyProperty{}))
		base.properties.Add(operation.Property(&CommandFlagsProperty{}))
		base.properties.Add(operation.Property(&CommandOutputProperty{}))
		base.properties.Add(operation.Property(&CommandErrorProperty{}))
		base.properties.Add(operation.Property(&CommandInputProperty{}))
		base.properties.Add(operation.Property(&CommandContextProperty{}))
	}
	return base.properties
}
