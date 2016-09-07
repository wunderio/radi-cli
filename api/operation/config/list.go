package config

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * An operation that lists configuration keys
 */

// Base class for config list Operation
type BaseConfigListOperation struct{}

// Id the operation
func (list *BaseConfigListOperation) Id() string {
	return "config.list"
}

// Label the operation
func (list *BaseConfigListOperation) Label() string {
	return "Config List"
}

// Description for the operation
func (list *BaseConfigListOperation) Description() string {
	return "List keyed configurations."
}

// Is this an internal API operation
func (list *BaseConfigListOperation) Internal() bool {
	return false
}
func (list *BaseConfigListOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}
