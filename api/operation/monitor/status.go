package monitor

const (
	OPERATION_ID_MONITOR_STATUS = "monitor.status"
)

/**
 * Status Operations are meant to return short parseable status
 * messages, about the state of the app
 */

// Base class for monitor status Operation
type BaseMonitorStatusOperation struct {
	MonitorBaseWriterOperation
}

// Id the operation
func (status *BaseMonitorStatusOperation) Id() string {
	return OPERATION_ID_MONITOR_STATUS
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
