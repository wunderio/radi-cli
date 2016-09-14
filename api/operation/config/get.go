package config

import (
	"errors"

	log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

const (
	OPERATION_ID_CONFIG_GET = "config.get"
)

/**
 * Retrieve keyed configurations for the config handler
 */

// Base class for config get Operation
type BaseConfigGetOperation struct {
	BaseConfigKeyValueOperation
}

// Id the operation
func (get *BaseConfigGetOperation) Id() string {
	return OPERATION_ID_CONFIG_GET
}

// Label the operation
func (get *BaseConfigGetOperation) Label() string {
	return "Config Get"
}

// Description for the operation
func (get *BaseConfigGetOperation) Description() string {
	return "Retrieve a keyed configuration."
}

// Is this an internal API operation
func (get *BaseConfigGetOperation) Internal() bool {
	return false
}

//
type ConfigConnectorGetOperation struct {
	BaseConfigGetOperation
	BaseConfigConnectorOperation
}

//
func (get ConfigConnectorGetOperation) Validate() bool {
	return true
}

//
func (get ConfigConnectorGetOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	confs := get.BaseConfigKeyValueOperation.Configurations()
	keyConf, _ := confs.Get(OPERATION_CONFIGURATION_CONFIG_KEY)
	valueConf, _ := confs.Get(OPERATION_CONFIGURATION_CONFIG_VALUE)

	if key, ok := keyConf.Get().(string); ok {
		if value, ok := get.Connector().Get(key); ok {
			valueConf.Set(value)
		} else {
			log.Error("Config connector did not find the value you were looking for")
			result.Set(false, []error{errors.New("Config connector did not find the value you were looking for")})
		}
	} else {
		log.Error("Could not get a string value for Key from the config connector")
		result.Set(false, []error{errors.New("Could not get a string value for Key from the config connector")})
	}

	return operation.Result(&result)
}
