package config

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Set keyed values into the Config handler
 */

// Base class for config set Operation
type BaseConfigSetOperation struct{}

// Id the operation
func (set *BaseConfigSetOperation) Id() string {
	return "config.set"
}

// Label the operation
func (set *BaseConfigSetOperation) Label() string {
	return "Config Set"
}

// Description for the operation
func (set *BaseConfigSetOperation) Description() string {
	return "Set a keyed configuration."
}

// Is this an internal API operation
func (set *BaseConfigSetOperation) Internal() bool {
	return false
}
func (set *BaseConfigSetOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}
