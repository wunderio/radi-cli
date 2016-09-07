package monitor

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Status Operations are meant to return short parseable status
 * messages, about the state of the app
 */

// Base class for monitor status Operation
type BaseMonitorStatusOperation struct {}
// Id the operation
func (status *BaseMonitorStatusOperation) Id() string {
	return "monitor.status"
}
// Label the operation
func (status *BaseMonitorStatusOperation) Label() string {
	return "Status"
}
// Description for the operation
func (status *BaseMonitorStatusOperation) Description() string {
	return "App status information."
}
// Is this an internal API operation
func (status *BaseMonitorStatusOperation) Internal() bool {
	return false
}
func (status *BaseMonitorStatusOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}
