package config

// Provide a Connector to Config sources
type ConfigConnector interface {
	Readers(key string) ScopedReaders
	Writers(key string) ScopedWriters
	List() []string
}

/**
 * Base operations for configs, primarily giving accesors for the connector
 */

// Constructor for BaseConfigConnectorOperation
func New_BaseConfigConnectorOperation(connector ConfigConnector) *BaseConfigConnectorOperation {
	return &BaseConfigConnectorOperation{
		connector: connector,
	}
}

// A Base config operation that provides a config connector
type BaseConfigConnectorOperation struct {
	connector ConfigConnector
}

// set the operation config connect
func (base *BaseConfigConnectorOperation) SetConnector(connector ConfigConnector) {
	base.connector = connector
}

// retrieve the operations config connnector
func (base *BaseConfigConnectorOperation) Connector() ConfigConnector {
	return base.connector
}
