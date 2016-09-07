package monitor

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Status Operations are meant to return long human readable
 * messages, about the state of the app
 */

// Base class for monitor info Operation
type BaseMonitorInfoOperation struct {}
// Id the operation
func (info *BaseMonitorInfoOperation) Id() string {
	return "monitor.info"
}
// Label the operation
func (info *BaseMonitorInfoOperation) Label() string {
	return "Information"
}
// Description for the operation
func (info *BaseMonitorInfoOperation) Description() string {
	return "App information."
}
// Is this an internal API operation
func (info *BaseMonitorInfoOperation) Internal() bool {
	return false
}
func (info *BaseMonitorInfoOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}
