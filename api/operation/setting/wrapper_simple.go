package setting

import (
	"errors"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * A simplified Settings wrapper, that performs blocking operations
 * using the operation.Operations.  This can be used to simplify
 * using the settings operations as a single struct.
 */

// Constructor for SimpleSettingWrapper
func New_SimpleSettingWrapper(operations *operation.Operations) *SimpleSettingWrapper {
	return &SimpleSettingWrapper{
		operations: operations,
	}
}

// A simple, blocking, Settings operations wrapper
type SimpleSettingWrapper struct {
	operations *operation.Operations
}

func (wrapper *SimpleSettingWrapper) Get(key string) (string, error) {
	var found, success bool
	var op operation.Operation
	var keyProp, valueProp operation.Property
	var err []error

	result := ""

	if op, found = wrapper.operations.Get(OPERATION_ID_SETTING_GET); !found {
		return result, errors.New("No get operation available in Setting Wrapper")
	}

	if keyProp, found = op.Properties().Get(OPERATION_PROPERTY_SETTING_KEY); !found {
		return result, errors.New("No key property available in Get operation in Setting Wrapper")
	}

	if !keyProp.Set(key) {
		return result, errors.New("Key property value failed to set in Setting Wrapper")
	}

	if valueProp, found = op.Properties().Get(OPERATION_PROPERTY_SETTING_VALUE); !found {
		return result, errors.New("No value property available in Get operation in Setting Wrapper")
	}

	if success, err = op.Exec().Success(); !success {
		return result, err[0] //errors.New("Operation get failed to execute in Setting Wrapper")
	}

	result = string(valueProp.Get().([]byte))

	return result, nil
}
func (wrapper *SimpleSettingWrapper) Set(key, value string) error {
	var found, success bool
	var op operation.Operation
	var keyProp, valueProp operation.Property
	var err []error

	if op, found = wrapper.operations.Get(OPERATION_ID_SETTING_SET); !found {
		return errors.New("No set operation available in Setting Wrapper")
	}

	if keyProp, found = op.Properties().Get(OPERATION_PROPERTY_SETTING_KEY); !found {
		return errors.New("No key property available in Set operation in Setting Wrapper")
	}

	if valueProp, found = op.Properties().Get(OPERATION_PROPERTY_SETTING_VALUE); !found {
		return errors.New("No value property available in Set operation in Setting Wrapper")
	}

	if !keyProp.Set(key) {
		return errors.New("Key property failed to set in Setting Wrapper")
	}

	if !valueProp.Set([]byte(value)) {
		return errors.New("Value property failed to set in Setting Wrapper")
	}

	if success, err = op.Exec().Success(); !success {
		return err[0] //errors.New("Operation get failed to execute in Setting Wrapper")
	}

	return nil
}
func (wrapper *SimpleSettingWrapper) List(parent string) ([]string, error) {
	var found, success bool
	var op operation.Operation
	var keyProp, keysProp operation.Property
	var err []error

	result := []string{}

	if op, found = wrapper.operations.Get(OPERATION_ID_SETTING_LIST); !found {
		return result, errors.New("No list operation available in Setting Wrapper")
	}

	if keyProp, found = op.Properties().Get(OPERATION_PROPERTY_SETTING_KEY); !found {
		return result, errors.New("No Parent key property available in Get operation in Setting Wrapper")
	}

	if !keyProp.Set(parent) {
		return result, errors.New("Parent key property value failed to set in Setting Wrapper")
	}

	if keysProp, found = op.Properties().Get(OPERATION_PROPERTY_SETTING_KEYS); !found {
		return result, errors.New("No keys value property available in list operation in Setting Wrapper")
	}

	if success, err = op.Exec().Success(); !success {
		return result, err[0] //errors.New("Operation get failed to execute in Setting Wrapper")
	}

	result = keysProp.Get().([]string)

	return result, nil
}
