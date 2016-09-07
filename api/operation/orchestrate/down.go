package orchestrate

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Orchestration DOWN - like docker-compose down
 *
 * Bring down all app containers, and remove them an all
 * related volumes and networks.
 */

// Base class for orchestration Down Operation
type BaseOrchestrationDownOperation struct{}

// Id the operation
func (down *BaseOrchestrationDownOperation) Id() string {
	return "orchestrate.down"
}

// Label the operation
func (down *BaseOrchestrationDownOperation) Label() string {
	return "Down"
}

// Description for the operation
func (down *BaseOrchestrationDownOperation) Description() string {
	return "This operation will bring down all containers, volumes and networks related to an application."
}

// Is this an internal API operation
func (down *BaseOrchestrationDownOperation) Internal() bool {
	return false
}
func (down *BaseOrchestrationDownOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}
