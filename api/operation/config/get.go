package config

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Retrieve keyed configurations for the config handler
 */

// Base class for config get Operation
type BaseConfigGetOperation struct{}

// Id the operation
func (get *BaseConfigGetOperation) Id() string {
	return "config.get"
}

// Label the operation
func (get *BaseConfigGetOperation) Label() string {
	return "Config Get"
}

// Description for the operation
func (get *BaseConfigGetOperation) Description() string {
	return "Retrieve a keyed configuration."
}

// Is this an internal API operation
func (get *BaseConfigGetOperation) Internal() bool {
	return false
}
func (get *BaseConfigGetOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}
