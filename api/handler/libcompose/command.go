package libcompose

import (
	"errors"

	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/command"
)

/**
 * Implement command containers that mix into
 * libCompose orchestrated containers
 */

const (
	CONFIG_KEY_COMMAND = "commands" // The Config key for settings
)

// A wrapper interface which pulls command information from a config wrapper backend
type CommandConfigWrapper interface {
	List(parent string) ([]string, error)
	Get(key string) (*CommandYmlCommand, error)
}

/**
 * Operations
 */

// LibCompose Command List operation
type LibcomposeCommandListOperation struct {
	command.BaseCommandListOperation
	command.BaseCommandKeyKeysOperation
	BaseLibcomposeNameFilesOperation

	Wrapper    CommandConfigWrapper
	properties *operation.Properties
}

// Validate the operation
func (list *LibcomposeCommandListOperation) Validate() bool {
	return true
}

// Get properties
func (list *LibcomposeCommandListOperation) Properties() *operation.Properties {
	baseProps := list.BaseCommandListOperation.Properties()

	keyKeysProps := list.BaseCommandKeyKeysOperation.Properties()
	baseProps.Merge(*keyKeysProps)

	return baseProps
}

// Execute the libCompose Command List operation
func (list *LibcomposeCommandListOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	props := list.BaseCommandKeyKeysOperation.Properties()
	keyProp, _ := props.Get(command.OPERATION_PROPERTY_COMMAND_KEY)
	keysProp, _ := props.Get(command.OPERATION_PROPERTY_COMMAND_KEYS)

	parent := ""
	if key, ok := keyProp.Get().(string); ok && key != "" {
		parent = key
	}

	if keyList, err := list.Wrapper.List(parent); err == nil {
		keysProp.Set(keyList)
	} else {
		result.Set(false, []error{err})
	}

	return operation.Result(&result)
}

// LibCompose Command Get operation
type LibcomposeCommandGetOperation struct {
	command.BaseCommandGetOperation
	command.BaseCommandKeyCommandOperation
	BaseLibcomposeNameFilesOperation

	Wrapper    CommandConfigWrapper
	properties *operation.Properties
}

// Validate the operation
func (get *LibcomposeCommandGetOperation) Validate() bool {
	return true
}

// Get properties
func (get *LibcomposeCommandGetOperation) Properties() *operation.Properties {
	baseProps := get.BaseCommandGetOperation.Properties()

	keyCommandProps := get.BaseCommandKeyCommandOperation.Properties()
	baseProps.Merge(*keyCommandProps)

	return baseProps
}

// Execute the libCompose Command Get operation
func (get *LibcomposeCommandGetOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	props := get.BaseCommandKeyCommandOperation.Properties()
	keyProp, _ := props.Get(command.OPERATION_PROPERTY_COMMAND_KEY)
	commandProp, _ := props.Get(command.OPERATION_PROPERTY_COMMAND_COMMAND)

	if key, ok := keyProp.Get().(string); ok && key != "" {

		if comYml, err := get.Wrapper.Get(key); err == nil {
			// pass all props to make a project
			project, _ := MakeComposeProject(get.BaseLibcomposeNameFilesOperation.Properties())
			com := comYml.Command(project)
			commandProp.Set(com)
		} else {
			result.Set(false, []error{err})
		}

	} else {
		result.Set(false, []error{errors.New("No command name provided.")})
	}

	return operation.Result(&result)
}
