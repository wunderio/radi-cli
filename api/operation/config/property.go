package config

import (
	log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Here are the commond shared propertys for the various
 * Config operations.
 */

const (
	// config for a single config key
	OPERATION_PROPERTY_CONFIG_KEY = "config.key"
	// config for an orerered list of keys
	OPERATION_PROPERTY_CONFIG_KEYS = "config.keys"
	// config for a single config value ([]byte])
	OPERATION_PROPERTY_CONFIG_VALUE = "config.value"
	// config for a single config value ([]byte])
	OPERATION_PROPERTY_CONFIG_SCOPE = "config.scope"
	// config for a single config value (as an io.readet)
	OPERATION_PROPERTY_CONFIG_VALUE_READERS = "config.value.reader"
	// config for a single config value (as an io.writer)
	OPERATION_PROPERTY_CONFIG_VALUE_WRITERS = "config.value.writer"
)

// property for a single config ket
type ConfigKeyProperty struct {
	operation.StringProperty
}

// Id for the property
func (confKey *ConfigKeyProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_KEY
}

// Label for the property
func (confKey *ConfigKeyProperty) Label() string {
	return "property key."
}

// Description for the property
func (confKey *ConfigKeyProperty) Description() string {
	return "property key."
}

// property for an ordered list of config keys
type ConfigKeysProperty struct {
	operation.StringSliceProperty
}

// Id for the property
func (keyValue *ConfigKeysProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_KEYS
}

// Label for the property
func (keyValue *ConfigKeysProperty) Label() string {
	return "property key list."
}

// Description for the property
func (keyValue *ConfigKeysProperty) Description() string {
	return "property key list."
}

// property for a single config value
type ConfigValueProperty struct {
	operation.BytesArrayProperty
}

// Id for the property
func (confValue *ConfigValueProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_VALUE
}

// Label for the property
func (confValue *ConfigValueProperty) Label() string {
	return "property value."
}

// Description for the property
func (confValue *ConfigValueProperty) Description() string {
	return "property value."
}

// property for a value as a set of io.Readers
type ConfigValueScopedReadersProperty struct {
	value ScopedReaders
}

// Id for the property
func (confValue *ConfigValueScopedReadersProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_VALUE_READERS
}

// Label for the property
func (confValue *ConfigValueScopedReadersProperty) Label() string {
	return "Config value readers."
}

// Description for the property
func (confValue *ConfigValueScopedReadersProperty) Description() string {
	return "Config value in the form of an ScopeReaders, which is an ordered map of io.Readers."
}

// Retreive the property value
func (property *ConfigValueScopedReadersProperty) Get() interface{} {
	return interface{}(property.value)
}

// Assign the property value
func (property *ConfigValueScopedReadersProperty) Set(value interface{}) bool {
	if converted, ok := value.(ScopedReaders); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected ScopedReaders")
		return false
	}
}

// property for a single value as an io.Writer
type ConfigValueScopedWritersProperty struct {
	value ScopedWriters
}

// Id for the property
func (confValue *ConfigValueScopedWritersProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_VALUE_WRITERS
}

// Label for the property
func (confValue *ConfigValueScopedWritersProperty) Label() string {
	return "Config value writers."
}

// Description for the property
func (confValue *ConfigValueScopedWritersProperty) Description() string {
	return "Config value in the form of an io.Writer."
}

// Retreive the property value
func (property *ConfigValueScopedWritersProperty) Get() interface{} {
	return interface{}(property.value)
}

// Assign the property value
func (property *ConfigValueScopedWritersProperty) Set(value interface{}) bool {
	if converted, ok := value.(ScopedWriters); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected ScopedWriters")
		return false
	}
}
