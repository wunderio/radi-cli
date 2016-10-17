package config

import (
	"errors"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * The Simple config connector operations are config operations that use
 * a connector is the simplest way, by directly calling the operation and
 * waiting for a response.
 *
 * @NOTE these operations can result in blocking of the connector stalls
 *   or fails.  This means that either we rely on an advanced connector
 *   which can itself time out, or we write more advanced operations.
 */

// Config Get operation that relies on a ConfigConnector for an io.Reader
type ConfigSimpleConnectorReadersOperation struct {
	BaseConfigReadersOperation
	BaseConfigKeyReadersOperation
	BaseConfigConnectorOperation
}

// Validate the operation
func (readers ConfigSimpleConnectorReadersOperation) Validate() bool {
	return true
}

// Execute the operation
func (readers ConfigSimpleConnectorReadersOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	props := readers.BaseConfigKeyReadersOperation.Properties()
	keyProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_KEY)
	readersProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_VALUE_READERS)

	if key, ok := keyProp.Get().(string); ok && key != "" {
		if readersValue := readers.Connector().Readers(key); len(readersValue.Order()) > 0 {
			readersProp.Set(readersValue)
		} else {
			result.Set(false, []error{errors.New("Unknown config key requested")})
		}
	} else {
		result.Set(false, []error{errors.New("Invalid config key requested")})
	}

	return operation.Result(&result)
}

// Config Set operation that relies on a ConfigConnector for an io.Writer
type ConfigSimpleConnectorWritersOperation struct {
	BaseConfigWritersOperation
	BaseConfigKeyWritersOperation
	BaseConfigConnectorOperation
}

// Validate the operation
func (writers ConfigSimpleConnectorWritersOperation) Validate() bool {
	return true
}

// Execute the operation
func (writers ConfigSimpleConnectorWritersOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	props := writers.BaseConfigKeyWritersOperation.Properties()
	keyProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_KEY)
	writersProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_VALUE_WRITERS)

	if key, ok := keyProp.Get().(string); ok && key != "" {
		if writersValue := writers.Connector().Writers(key); len(writersValue.Order()) > 0 {
			writersProp.Set(writersValue)
		} else {
			result.Set(false, []error{errors.New("Unknown config key requested")})
		}
	} else {
		result.Set(false, []error{errors.New("Invalid config key for config writers")})
	}

	return operation.Result(&result)
}

// Config List operation that relies on a ConfigConnector for an io.Writer
type ConfigSimpleConnectorListOperation struct {
	BaseConfigListOperation
	BaseConfigKeyKeysOperation
	BaseConfigConnectorOperation
}

// Validate the operation
func (list ConfigSimpleConnectorListOperation) Validate() bool {
	return true
}

// Execute the operation
func (list ConfigSimpleConnectorListOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	props := list.BaseConfigKeyKeysOperation.Properties()
	keyProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_KEY)
	keysProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_KEYS)

	if key, ok := keyProp.Get().(string); ok || key == "" {
		if list := list.Connector().List(); len(list) > 0 {
			keysProp.Set(list)
		} else {
			result.Set(false, []error{errors.New("Config has no keys")})
		}
	} else {
		result.Set(false, []error{errors.New("Invalid config parent key provided for config list")})
	}

	return operation.Result(&result)
}
