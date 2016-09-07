package security

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Operations for authenticating access to the API
 * which can come in a few forms.  This files holds
 * the Base operations, on which handlers should
 * build.
 */

// Base class for security authenticate Operation
type BaseSecurityAuthenticateOperation struct{}

// Id the operation
func (authenticate *BaseSecurityAuthenticateOperation) Id() string {
	return "security.authenticate"
}

// Label the operation
func (authenticate *BaseSecurityAuthenticateOperation) Label() string {
	return "Authenticate"
}

// Description for the operation
func (authenticate *BaseSecurityAuthenticateOperation) Description() string {
	return "Authenticate access to the app."
}

// Is this an internal API operation
func (authenticate *BaseSecurityAuthenticateOperation) Internal() bool {
	return false
}
func (authenticate *BaseSecurityAuthenticateOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}
