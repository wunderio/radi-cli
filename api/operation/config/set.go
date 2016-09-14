package config

import (
	"errors"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

const (
	OPERATION_ID_CONFIG_SET = "config.set"
)

/**
 * Set keyed values into the Config handler
 */

// Base class for config set Operation
type BaseConfigSetOperation struct {
	BaseConfigKeyValueOperation
}

// Id the operation
func (set *BaseConfigSetOperation) Id() string {
	return OPERATION_ID_CONFIG_SET
}

// Label the operation
func (set *BaseConfigSetOperation) Label() string {
	return "Config Set"
}

// Description for the operation
func (set *BaseConfigSetOperation) Description() string {
	return "Set a keyed configuration."
}

// Is this an internal API operation
func (set *BaseConfigSetOperation) Internal() bool {
	return false
}

//
type ConfigConnectorSetOperation struct {
	BaseConfigSetOperation
	BaseConfigConnectorOperation
}

//
func (set ConfigConnectorSetOperation) Validate() bool {
	return true
}

//
func (set ConfigConnectorSetOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	confs := set.Configurations()
	keyConf, _ := confs.Get(OPERATION_CONFIGURATION_CONFIG_KEY)
	valueConf, _ := confs.Get(OPERATION_CONFIGURATION_CONFIG_VALUE)

	if key, okKey := keyConf.Get().(string); okKey {
		if value, okValue := valueConf.Get().(string); okValue {
			if okSet := set.Connector().Set(key, value); !okSet {
				result.Set(false, []error{errors.New("")})
			}
		} else {
			result.Set(false, []error{errors.New("Could not retrieve Value conf for config set operation")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not assign value to key conf for config set operation")})
	}

	return operation.Result(&result)
}
