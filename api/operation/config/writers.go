package config

const (
	OPERATION_ID_CONFIG_WRITERS = "config.writers"
)

/**
 * Retrieve writers for the Config handler
 */

// Base class for config retrieve writers Operation
type BaseConfigWritersOperation struct {
	BaseConfigKeyWritersOperation
}

// Id the operation
func (writers *BaseConfigKeyWritersOperation) Id() string {
	return OPERATION_ID_CONFIG_WRITERS
}

// Label the operation
func (writers *BaseConfigKeyWritersOperation) Label() string {
	return "Config Writers"
}

// Description for the operation
func (writers *BaseConfigKeyWritersOperation) Description() string {
	return "Get a set of scoped writers for a configuration."
}

// Is this an internal API operation
func (writers *BaseConfigKeyWritersOperation) Internal() bool {
	return true
}
