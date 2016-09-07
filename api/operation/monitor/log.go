package monitor

import (
	log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Write messages to the log
 */

const (
	OPERATION_ID_MONITOR_LOG                         = "monitor.log"
	OPERATION_CONFIGURATION_CONF_MONITOR_LOG_TYPE    = "monitor.log.type"
	OPERATION_CONFIGURATION_CONF_MONITOR_LOG_MESSAGE = "monitor.log.message"
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
	configurations *operation.Configurations
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

// Add a Message configuration
func (logger *MonitorStandardLogOperation) Configurations() *operation.Configurations {
	if logger.configurations == nil {
		logger.configurations = &operation.Configurations{}

		logger.configurations.Add(operation.Configuration(NewMonitorLogTypeConfiguration("info")))
		logger.configurations.Add(operation.Configuration(&MonitorLogMessageConfiguration{}))
	}
	return logger.configurations
}

// Exec the log output
func (logger *MonitorStandardLogOperation) Exec() operation.Result {
	baseResult := operation.BaseResult{}

	// we ignore the conf tests, as we ensured that the conf would exist in the Configuration() method
	logTypeConf, _ := logger.Configurations().Get(OPERATION_CONFIGURATION_CONF_MONITOR_LOG_TYPE)
	messageConf, _ := logger.Configurations().Get(OPERATION_CONFIGURATION_CONF_MONITOR_LOG_MESSAGE)

	var logType, message string
	var ok bool

	if message, ok = messageConf.Get().(string); !ok {
		log.Error("MonitorStandardLogOperation has no message assigned")
	} else {
		if logType, ok = logTypeConf.Get().(string); !ok {
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

func NewMonitorLogTypeConfiguration(logType string) *MonitorLogTypeConfiguration {
	conf := MonitorLogTypeConfiguration{}
	conf.Set(logType)
	return &conf
}

// Configuration for a monitoring log status : error|info
type MonitorLogTypeConfiguration struct {
	operation.StringConfiguration
}

// Id for the configuration
func (logType *MonitorLogTypeConfiguration) Id() string {
	return OPERATION_CONFIGURATION_CONF_MONITOR_LOG_TYPE
}

// Label for the configuration
func (logType *MonitorLogTypeConfiguration) Label() string {
	return "Message type."
}

// Description for the configuration
func (logType *MonitorLogTypeConfiguration) Description() string {
	return "Message type, which can be either info or error."
}

// Configuration for a monitoring log message
type MonitorLogMessageConfiguration struct {
	operation.StringConfiguration
}

// Id for the configuration
func (message *MonitorLogMessageConfiguration) Id() string {
	return OPERATION_CONFIGURATION_CONF_MONITOR_LOG_MESSAGE
}

// Label for the configuration
func (message *MonitorLogMessageConfiguration) Label() string {
	return "Message to be logged."
}

// Description for the configuration
func (message *MonitorLogMessageConfiguration) Description() string {
	return "Message which will be sent to the standard logger."
}
