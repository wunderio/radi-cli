package config

const (
	OPERATION_ID_CONFIG_LIST = "config.list"
)

/**
 * Retrieve keyed configurations for the config handler
 */

// Base class for config list Operation
type BaseConfigListOperation struct {
	BaseConfigKeyKeysOperation
}

// Id the operation
func (list *BaseConfigListOperation) Id() string {
	return OPERATION_ID_CONFIG_LIST
}

// Label the operation
func (list *BaseConfigListOperation) Label() string {
	return "Config List"
}

// Description for the operation
func (list *BaseConfigListOperation) Description() string {
	return "Retrieve a list of available configuration keys."
}

// Is this an internal API operation
func (list *BaseConfigListOperation) Internal() bool {
	return true
}
