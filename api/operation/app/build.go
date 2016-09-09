package api

/**
 * Provide an operations which will build an API object
 */

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Execute a command using the command handler
 */

// Base class for command build Operation
type BaseAppBuildOperation struct{}

// Id the operation
func (build *BaseAppBuildOperation) Id() string {
	return "app.build"
}

// Label the operation
func (build *BaseAppBuildOperation) Label() string {
	return "Build the APP"
}

// Description for the operation
func (build *BaseAppBuildOperation) Description() string {
	return "Build the current API by adding more handler/operations to it."
}

// Is this an internal API operation
func (build *BaseAppBuildOperation) Internal() bool {
	return true
}
func (build *BaseAppBuildOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}
