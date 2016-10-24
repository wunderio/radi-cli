package monitor

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

const (
	OPERATION_PROPERTY_ID_MONITOR_WRITER = "monitor.output.writer"
)

// Configuration for an outputter for monitoring
type MonitorOutputProperty struct {
	operation.WriterProperty
}

// Id for the configuration
func (output *MonitorOutputProperty) Id() string {
	return OPERATION_PROPERTY_ID_MONITOR_WRITER
}

// Label for the configuration
func (output *MonitorOutputProperty) Label() string {
	return "Output handler for the monitor"
}

// Description for the operation
func (output *MonitorOutputProperty) Description() string {
	return "Attach an io.Writer to the operation, and it will be used to capture the output.  By default, the output will go to log."
}

// Is the Property internal only
func (output *MonitorOutputProperty) Internal() bool {
	return false
}
