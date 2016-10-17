package security

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * User operations return information about the currently
 * authenticated user
 */

// Base class for security authenticate Operation
type BaseSecurityUserOperation struct{}

// Id the operation
func (authenticate *BaseSecurityUserOperation) Id() string {
	return "security.user"
}

// Label the operation
func (authenticate *BaseSecurityUserOperation) Label() string {
	return "Get User"
}

// Description for the operation
func (authenticate *BaseSecurityUserOperation) Description() string {
	return "Retrieve information about the current app user."
}

// Is this an internal API operation
func (authenticate *BaseSecurityUserOperation) Internal() bool {
	return false
}
func (authenticate *BaseSecurityUserOperation) Properties() *operation.Properties {
	return &operation.Properties{}
}
