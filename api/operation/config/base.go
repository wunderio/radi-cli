package config

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

//
type BaseConfigConnectorOperation struct {
	connector ConfigConnector
}

//
func (base *BaseConfigConnectorOperation) SetConnector(connector ConfigConnector) {
	base.connector = connector
}

//
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

		base.configurations.Add(operation.Configuration(&ConfigValueConfiguration{}))
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
type BaseConfigKeyValueROOperation struct {
	configurations *operation.Configurations
}

//
func (base *BaseConfigKeyValueROOperation) Configurations() *operation.Configurations {
	if base.configurations == nil {
		base.configurations = &operation.Configurations{}

		base.configurations.Add(operation.Configuration(&ConfigKeyConfiguration{}))
		base.configurations.Add(operation.Configuration(&ConfigValueROConfiguration{}))
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
