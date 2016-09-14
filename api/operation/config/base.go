package config

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

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

//
type BaseConfigKeyOperation struct {
	configurations *operation.Configurations
}

//
func (base *BaseConfigKeyOperation) Configurations() *operation.Configurations {
	if base.configurations == nil {
		base.configurations = &operation.Configurations{}

		base.configurations.Add(operation.Configuration(&ConfigKeyConfiguration{}))
	}
	return base.configurations
}

//
type BaseConfigKeyValueOperation struct {
	configurations *operation.Configurations
}

//
func (base *BaseConfigKeyValueOperation) Configurations() *operation.Configurations {
	if base.configurations == nil {
		base.configurations = &operation.Configurations{}

		base.configurations.Add(operation.Configuration(&ConfigKeyConfiguration{}))
		base.configurations.Add(operation.Configuration(&ConfigValueConfiguration{}))
	}
	return base.configurations
}

//
type BaseConfigKeyKeysOperation struct {
	configurations *operation.Configurations
}

//
func (base *BaseConfigKeyKeysOperation) Configurations() *operation.Configurations {
	if base.configurations == nil {
		base.configurations = &operation.Configurations{}

		base.configurations.Add(operation.Configuration(&ConfigKeyConfiguration{}))
		base.configurations.Add(operation.Configuration(&ConfigKeysConfiguration{}))
	}
	return base.configurations
}
