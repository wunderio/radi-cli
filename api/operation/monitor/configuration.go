package monitor

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

const (
	OPERATION_CONFIGURATION_ID_MONITOR_WRITER = "monitor.output.writer"
)

// Configuration for an outputter for monitoring
type MonitorOutputConfiguration struct {
	operation.WriterConfiguration
}

// Id for the configuration
func (output *MonitorOutputConfiguration) Id() string {
	return "monitor.output.writer"
}

// Label for the configuration
func (output *MonitorOutputConfiguration) Label() string {
	return "Output handler for the monitor"
}

// Description for the configuration
func (output *MonitorOutputConfiguration) Description() string {
	return "Attach an io.Writer to the configuration, and it will be used to capture the output.  By default, the output will go to log."
}
