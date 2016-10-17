package monitor

import (
	"io"

	log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

// A handler base that writes to an outputter
type MonitorBaseWriterOperation struct{}

// A utility function to write a message to the configured writer
func (op *MonitorBaseWriterOperation) WriteMessage(message string) bool {
	if writerConfig, exists := op.Properties().Get(OPERATION_PROPERTY_ID_MONITOR_WRITER); exists {
		confValue := writerConfig.Get()
		if writer, ok := confValue.(io.Writer); ok {
			writer.Write([]byte(message))
			return true
		} else {
			log.WithFields(log.Fields{"writer": writer}).Warning("Could not write status, as the output configuration contains an invalid writer.")
		}
	}
	return false
}

// Add a writer configuration
func (op *MonitorBaseWriterOperation) Properties() *operation.Properties {
	configurations := operation.Properties{}

	configurations.Add(operation.Property(&MonitorOutputProperty{}))

	return &configurations
}
