package security

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Authorize operations should be used to question whether or not
 * the authorized user should be given access to other operations
 *
 * @QUESTION should authentication be tied to Operation.Id() values,
 * or should we also allow user|action|key style options
 */

// Base class for security authorize Operation
type BaseSecurityAuthorizeOperation struct{}

// Id the operation
func (authorize *BaseSecurityAuthorizeOperation) Id() string {
	return "security.authorize"
}

// Label the operation
func (authorize *BaseSecurityAuthorizeOperation) Label() string {
	return "Authorize"
}

// Description for the operation
func (authorize *BaseSecurityAuthorizeOperation) Description() string {
	return "Authorize access to a part of the app."
}

// Is this an internal API operation
func (authorize *BaseSecurityAuthorizeOperation) Internal() bool {
	return false
}
func (authorize *BaseSecurityAuthorizeOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}
