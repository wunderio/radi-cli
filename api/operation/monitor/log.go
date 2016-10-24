package monitor

import (
	log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Write messages to the log
 */

const (
	OPERATION_ID_MONITOR_LOG                    = "monitor.log"
	OPERATION_PROPERTY_CONF_MONITOR_LOG_TYPE    = "monitor.log.type"
	OPERATION_PROPERTY_CONF_MONITOR_LOG_MESSAGE = "monitor.log.message"
)

// Base class for monitor log Operation
type BaseMonitorLogOperation struct {
	MonitorBaseWriterOperation
}

// Id the operation
func (logger *BaseMonitorLogOperation) Id() string {
	return OPERATION_ID_MONITOR_LOG
}

// Label the operation
func (logger *BaseMonitorLogOperation) Label() string {
	return "Log a message"
}

// Description for the operation
func (logger *BaseMonitorLogOperation) Description() string {
	return "Log a message."
}

// Is this an internal API operation
func (logger *BaseMonitorLogOperation) Internal() bool {
	return false
}

// Is the log request valid
func (logger *BaseMonitorLogOperation) Validate() bool {
	return false
}

// Standard output logger
type MonitorStandardLogOperation struct {
	properties *operation.Properties
}

// Id the operation
func (logger *MonitorStandardLogOperation) Id() string {
	return OPERATION_ID_MONITOR_LOG
}

// Label the operation
func (logger *MonitorStandardLogOperation) Label() string {
	return "Log a message"
}

// Description for the operation
func (logger *MonitorStandardLogOperation) Description() string {
	return "Log a message."
}

// Is this an internal API operation
func (logger *MonitorStandardLogOperation) Internal() bool {
	return false
}

// Is the log request valid
func (logger *MonitorStandardLogOperation) Validate() bool {
	return false
}

// Add a Message propery
func (logger *MonitorStandardLogOperation) Properties() *operation.Properties {
	if logger.properties == nil {
		logger.properties = &operation.Properties{}

		logger.properties.Add(operation.Property(NewMonitorLogTypeProperty("info")))
		logger.properties.Add(operation.Property(&MonitorLogMessageProperty{}))
	}
	return logger.properties
}

// Exec the log output
func (logger *MonitorStandardLogOperation) Exec() operation.Result {
	baseResult := operation.BaseResult{}

	// we ignore the conf tests, as we ensured that the conf would exist in the property() method
	logTypeProp, _ := logger.Properties().Get(OPERATION_PROPERTY_CONF_MONITOR_LOG_TYPE)
	messageProp, _ := logger.Properties().Get(OPERATION_PROPERTY_CONF_MONITOR_LOG_MESSAGE)

	var logType, message string
	var ok bool

	if message, ok = messageProp.Get().(string); !ok {
		log.Error("MonitorStandardLogOperation has no message assigned")
	} else {
		if logType, ok = logTypeProp.Get().(string); !ok {
			logType = "info"
		}
		switch logType {
		case "error":
			log.Error(message)

		case "info":
			fallthrough
		default:
			log.Info(message)

		}
	}

	baseResult.Set(true, []error{})
	return operation.Result(&baseResult)
}

func NewMonitorLogTypeProperty(logType string) *MonitorLogTypeProperty {
	conf := MonitorLogTypeProperty{}
	conf.Set(logType)
	return &conf
}

// property for a monitoring log status : error|info
type MonitorLogTypeProperty struct {
	operation.StringProperty
}

// Id for the property
func (logType *MonitorLogTypeProperty) Id() string {
	return OPERATION_PROPERTY_CONF_MONITOR_LOG_TYPE
}

// Label for the property
func (logType *MonitorLogTypeProperty) Label() string {
	return "Message type."
}

// Description for the property
func (logType *MonitorLogTypeProperty) Description() string {
	return "Message type, which can be either info or error."
}

// Is the Property internal only
func (logType *MonitorLogTypeProperty) Internal() bool {
	return false
}

// property for a monitoring log message
type MonitorLogMessageProperty struct {
	operation.StringProperty
}

// Id for the property
func (message *MonitorLogMessageProperty) Id() string {
	return OPERATION_PROPERTY_CONF_MONITOR_LOG_MESSAGE
}

// Label for the property
func (message *MonitorLogMessageProperty) Label() string {
	return "Message to be logged."
}

// Description for the property
func (message *MonitorLogMessageProperty) Description() string {
	return "Message which will be sent to the standard logger."
}

// Is the Property internal only
func (message *MonitorLogMessageProperty) Internal() bool {
	return false
}
