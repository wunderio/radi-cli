package config

const (
	OPERATION_ID_CONFIG_READERS = "config.readers"
)

/**
 * Retrieve keyed configurations for the config handler
 */

// Base class for config get Operation
type BaseConfigReadersOperation struct {
	BaseConfigKeyReadersOperation
}

// Id the operation
func (readers *BaseConfigReadersOperation) Id() string {
	return OPERATION_ID_CONFIG_READERS
}

// Label the operation
func (readers *BaseConfigReadersOperation) Label() string {
	return "Config Readers"
}

// Description for the operation
func (readers *BaseConfigReadersOperation) Description() string {
	return "Retrieve a keyed configuration scoped reader set."
}

// Is this an internal API operation
func (readers *BaseConfigReadersOperation) Internal() bool {
	return true
}
