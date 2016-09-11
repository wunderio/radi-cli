package config

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

const (
	OPERATION_ID_CONFIG_LIST = "config.list"
)

/**
 * Retrieve keyed configurations for the config handler
 */

// Base class for config list Operation
type BaseConfigListOperation struct {
	BaseConfigKeyKeysOperation
}

// Id the operation
func (list *BaseConfigListOperation) Id() string {
	return OPERATION_ID_CONFIG_LIST
}

// Label the operation
func (list *BaseConfigListOperation) Label() string {
	return "Config Get"
}

// Description for the operation
func (list *BaseConfigListOperation) Description() string {
	return "Retrieve a keyed configuration."
}

// Is this an internal API operation
func (list *BaseConfigListOperation) Internal() bool {
	return false
}

//
type ConfigConnectorListOperation struct {
	BaseConfigListOperation
	BaseConfigConnectorOperation
}

//
func (list ConfigConnectorListOperation) Validate() bool {
	return true
}

//
func (list ConfigConnectorListOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	confs := list.BaseConfigKeyKeysOperation.Configurations()
	keyConf, _ := confs.Get(OPERATION_CONFIGURATION_CONFIG_KEY)
	keysConf, _ := confs.Get(OPERATION_CONFIGURATION_CONFIG_KEYS)

	if key, ok := keyConf.Get().(string); ok && key != "" {
		keysConf.Set(list.Connector().List(key))
	} else {
		keysConf.Set(list.Connector().List(""))
	}

	return operation.Result(&result)
}
