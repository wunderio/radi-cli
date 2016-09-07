package orchestrate

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Orchestration UP - like docker-compose up
 *
 * Bring up all app containers, volumes and networks.
 */

// Base class for orchestration Up Operation
type BaseOrchestrationUpOperation struct{}

// Id the operation
func (up *BaseOrchestrationUpOperation) Id() string {
	return "orchestrate.up"
}

// Label the operation
func (up *BaseOrchestrationUpOperation) Label() string {
	return "Up"
}

// Description for the operation
func (up *BaseOrchestrationUpOperation) Description() string {
	return "This operation will bring up all containers, volumes and networks related to an application."
}

// Is this an internal API operation
func (up *BaseOrchestrationUpOperation) Internal() bool {
	return false
}
func (up *BaseOrchestrationUpOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}
