package command

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Base Operation classes for command implementation
 */

// A Base command operation that provides a single command key string and a command object
type BaseCommandKeyCommandOperation struct {
	properties *operation.Properties
}

// Return a static keys list Property
func (base *BaseCommandKeyCommandOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&CommandKeyProperty{}))
		base.properties.Add(operation.Property(&CommandCommandProperty{}))
	}
	return base.properties
}

// A Base command operation that returns a list of keys
type BaseCommandKeyKeysOperation struct {
	properties *operation.Properties
}

// Return a static keys list Property
func (base *BaseCommandKeyKeysOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&CommandKeyProperty{}))
		base.properties.Add(operation.Property(&CommandKeysProperty{}))
	}
	return base.properties
}

// A base command operation that provides a a key, flags strings, and various input and output Properties
type BaseCommandContextOperation struct {
	properties *operation.Properties
}

// get static Properties
func (base *BaseCommandContextOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&CommandContextProperty{}))
	}
	return base.properties
}

// A base command operation that provides a a key, flags strings, and various input and output Properties
type BaseCommandInputOutputOperation struct {
	properties *operation.Properties
}

// get static Properties
func (base *BaseCommandInputOutputOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&CommandInputProperty{}))
		base.properties.Add(operation.Property(&CommandOutputProperty{}))
		base.properties.Add(operation.Property(&CommandErrorProperty{}))
	}
	return base.properties
}

// A base command operation that provides akey, flags
type BaseCommandFlagsOperation struct {
	properties *operation.Properties
}

// get static Properties
func (base *BaseCommandFlagsOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&CommandFlagsProperty{}))
	}
	return base.properties
}
