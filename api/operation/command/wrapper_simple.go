package command

import (
	"errors"

	// log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

func New_SimpleCommandWrapper(operations *operation.Operations) *SimpleCommandWrapper {
	return &SimpleCommandWrapper{
		operations: operations,
	}
}

type SimpleCommandWrapper struct {
	operations *operation.Operations
}

func (wrapper *SimpleCommandWrapper) Get(key string) (Command, error) {
	var found, success bool
	var op operation.Operation
	var keyProp, commandProp operation.Property
	var err []error

	var result Command

	if op, found = wrapper.operations.Get(OPERATION_ID_COMMAND_GET); !found {
		return result, errors.New("No get operation available in Command Wrapper")
	}

	if keyProp, found = op.Properties().Get(OPERATION_PROPERTY_COMMAND_KEY); !found {
		return result, errors.New("No key property available in Get operation in Command Wrapper")
	}

	if !keyProp.Set(key) {
		return result, errors.New("Key property value failed to set in Command Wrapper")
	}

	if commandProp, found = op.Properties().Get(OPERATION_PROPERTY_COMMAND_COMMAND); !found {
		return result, errors.New("No command property available in Get operation in Command Wrapper")
	}

	if success, err = op.Exec().Success(); !success {
		return result, err[0] //errors.New("Operation get failed to execute in Command Wrapper")
	}

	result = commandProp.Get().(Command)

	return result, nil
}

func (wrapper *SimpleCommandWrapper) List(parent string) ([]string, error) {
	var found, success bool
	var op operation.Operation
	var keyProp, keysProp operation.Property
	var errs []error

	result := []string{}

	if op, found = wrapper.operations.Get(OPERATION_ID_COMMAND_LIST); !found {
		return result, errors.New("No list operation available in Command Wrapper")
	}

	if keyProp, found = op.Properties().Get(OPERATION_PROPERTY_COMMAND_KEY); !found {
		return result, errors.New("No key property available in Command Wrapper")
	}

	if !keyProp.Set(parent) {
		return result, errors.New("Key property value failed to set in Command Wrapper")
	}

	if keysProp, found = op.Properties().Get(OPERATION_PROPERTY_COMMAND_KEYS); !found {
		return result, errors.New("No keys property available in Command Wrapper")
	}

	if success, errs = op.Exec().Success(); !success {
		return result, errs[0]
	}

	result = keysProp.Get().([]string)
	return result, nil
}
