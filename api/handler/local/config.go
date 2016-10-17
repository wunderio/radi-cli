package local

import (
	"github.com/james-nesbitt/wundertools-go/api/handler/configconnect"
	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/config"
)

// A handler for local config
type LocalHandler_Config struct {
	LocalHandler_Base
}

// Identify the handler
func (handler *LocalHandler_Config) Id() string {
	return "local.config"
}

// [Handler.]Init tells the LocalHandler_Orchestrate to prepare it's operations
func (handler *LocalHandler_Config) Init() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	ops := operation.Operations{}

	// build a ConfigConnector for use with the Config operations.
	connector := configconnect.New_ConfigConnectYmlFiles(handler.settings.ConfigPaths)

	// Build this base operation to be shared across all of our config operations
	baseConnectorOperation := config.New_BaseConfigConnectorOperation(connector)

	// Now we can add config operations that use that Base class
	ops.Add(operation.Operation(&config.ConfigSimpleConnectorReadersOperation{BaseConfigConnectorOperation: *baseConnectorOperation}))
	ops.Add(operation.Operation(&config.ConfigSimpleConnectorWritersOperation{BaseConfigConnectorOperation: *baseConnectorOperation}))
	ops.Add(operation.Operation(&config.ConfigSimpleConnectorListOperation{BaseConfigConnectorOperation: *baseConnectorOperation}))

	handler.operations = &ops

	return operation.Result(&result)
}

// Make ConfigWrapper
func (handler *LocalHandler_Config) ConfigWrapper() config.ConfigWrapper {
	return config.ConfigWrapper(config.New_SimpleConfigWrapper(handler.operations))
}
