package local

import (
	"errors"

	"github.com/james-nesbitt/wundertools-go/api/handler/bytesource"
	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/config"
)

/**
 * Local handler provides operations based entirely on the
 * Local environment, primarily based on config files in
 * a project, based on the current path'
 */

// A handler that provides operations based on local project files
type LocalHandler struct {
	settings *LocalAPISettings

	configConnector config.ConfigConnector
}

// [Handler.]Id returns a string ID for the handler
func (handler *LocalHandler) Id() string {
	return "local"
}

// [Handler.]Init tells the LocalHandler to process it's configurations
func (handler *LocalHandler) Init() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	if handler.settings == nil {
		result.Set(false, []error{errors.New("No settings were passed into this API constructor.")})
	}

	if configConnector, err := handler.ConfigConnector(); err == nil {
		handler.configConnector = configConnector
	} else {
		result.Set(false, []error{err})
	}

	return operation.Result(&result)
}

// Get Local Operations
func (handler *LocalHandler) Operations() *operation.Operations {
	ops := operation.Operations{}

	baseConnectorOperation := config.BaseConfigConnectorOperation{}
	baseConnectorOperation.SetConnector(handler.configConnector)

	ops.Add(operation.Operation(&config.ConfigConnectorListOperation{BaseConfigConnectorOperation: baseConnectorOperation}))
	ops.Add(operation.Operation(&config.ConfigConnectorGetOperation{BaseConfigConnectorOperation: baseConnectorOperation}))
	ops.Add(operation.Operation(&config.ConfigConnectorSetOperation{BaseConfigConnectorOperation: baseConnectorOperation}))

	return &ops
}

// Make a Config Connector from the local config file, which can then be used for multiple operations
func (handler *LocalHandler) ConfigConnector() (config.ConfigConnector, error) {
	for _, key := range handler.settings.ConfigPaths.Order() {
		path, _ := handler.settings.ConfigPaths.Get(key)
		fileSource := path.FullPath(LOCAL_CONFIG_FILE)

		if fileSource.Validate() {
			ymlFileConfig := bytesource.NewYmlFileConfig_FromFileByteSource(fileSource)
			return config.ConfigConnector(ymlFileConfig), nil
		} else {
			return nil, errors.New("Filesource for config connector failed it's own validation.")
		}
	}

	return nil, errors.New("No valid config file found.")
}
