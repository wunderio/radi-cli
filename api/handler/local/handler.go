package local

import (
	"errors"
	"os"
	"path"

	"github.com/james-nesbitt/wundertools-go/api/handler/bytesource"
	"github.com/james-nesbitt/wundertools-go/api/handler/libcompose"
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
	settings        *LocalAPISettings
	configConnector config.ConfigConnector
	operations      *operation.Operations
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

	if configConnector, err := handler.getConfigConnector(); err == nil {
		handler.configConnector = configConnector
	} else {
		result.Set(false, []error{err})
	}

	// create new ops
	ops := operation.Operations{}
	handler.operations = &ops

	baseConfigConnectorOperation := config.BaseConfigConnectorOperation{}
	baseConfigConnectorOperation.SetConnector(handler.configConnector)

	ops.Add(operation.Operation(&config.ConfigConnectorListOperation{BaseConfigConnectorOperation: baseConfigConnectorOperation}))
	ops.Add(operation.Operation(&config.ConfigConnectorGetOperation{BaseConfigConnectorOperation: baseConfigConnectorOperation}))
	ops.Add(operation.Operation(&config.ConfigConnectorSetOperation{BaseConfigConnectorOperation: baseConfigConnectorOperation}))

	// now that we have added some config ops, let's use them to get settings when needed
	baseLibcomposeOrchestrate := handler.MakeBaseOrchestrationOperation(result)

	// Now we can add orchestration operations that use that Base class
	ops.Add(operation.Operation(&libcompose.LibcomposeMonitorLogsOperation{BaseLibcomposeOrchestrateNameFilesOperation: baseLibcomposeOrchestrate}))
	ops.Add(operation.Operation(&libcompose.LibcomposeOrchestrateUpOperation{BaseLibcomposeOrchestrateNameFilesOperation: baseLibcomposeOrchestrate}))
	ops.Add(operation.Operation(&libcompose.LibcomposeOrchestrateDownOperation{BaseLibcomposeOrchestrateNameFilesOperation: baseLibcomposeOrchestrate}))

	return operation.Result(&result)
}

// Get Local Operations
func (handler *LocalHandler) Operations() *operation.Operations {
	return handler.operations
}

// Make a Config Connector from the local config file, which can then be used for multiple operations
func (handler *LocalHandler) getConfigConnector() (config.ConfigConnector, error) {
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

// Run a config.get operation to retrieve a string value
func (handler *LocalHandler) configGet(key string) (string, bool) {
	if configGetOp, found := handler.operations.Get(config.OPERATION_ID_CONFIG_GET); found {
		// the handler contains an operation for config.get
		if keyConf, found := configGetOp.Configurations().Get(config.OPERATION_CONFIGURATION_CONFIG_KEY); found {
			// successfully found the config KEY conf
			if keyConf.Set(key) {
				// successfully set the key conf, so now we can exec()
				if success, _ := configGetOp.Exec().Success(); success {
					if valueConf, found := configGetOp.Configurations().Get(config.OPERATION_CONFIGURATION_CONFIG_VALUE); found {
						// successfully ran the operation
						return valueConf.Get().(string), true
					}
				}
			}
		}
	}
	// this handler failed to retrieve the setting
	return "", false
}

// A handoff function to make a base orchestration operation, which is really just a lot of linear code.
// @NOTE this needs the "config.get" operation to already be available
func (handler *LocalHandler) MakeBaseOrchestrationOperation(result operation.BaseResult) libcompose.BaseLibcomposeOrchestrateNameFilesOperation {
	// This Base operations will be at the root of all of the libCompose operations
	baseLibcomposeOrchestrate := libcompose.BaseLibcomposeOrchestrateNameFilesOperation{}
	orchestrateConfigurations := baseLibcomposeOrchestrate.Configurations()

	// Set a project name
	if projectName, found := handler.configGet("Project"); found {
		if projectNameConf, found := orchestrateConfigurations.Get(libcompose.OPERATION_CONFIGURATION_LIBCOMPOSE_PROJECTNAME); found {
			if !projectNameConf.Set(projectName) {
				result.Set(false, []error{errors.New("Could not set base libCompose project name.  Config set error on base Orchestration operation")})
			}
		} else {
			result.Set(false, []error{errors.New("Could not set base libCompose project name.  Config value not found on base Orchestration operation")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not set base libCompose project name.  Config value not found in handler config")})
	}

	// Add project docker-compose yml files
	if projectComposeFilesConf, found := orchestrateConfigurations.Get(libcompose.OPERATION_CONFIGURATION_LIBCOMPOSE_COMPOSEFILES); found {
		if !projectComposeFilesConf.Set([]string{path.Join(handler.settings.ProjectRootPath, "docker-compose.yml")}) {
			result.Set(false, []error{errors.New("Could not set base libcompose docker-compose file conf.  Config set error on base Orchestration operation")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not set base libcompose docker-compose file conf.  Config not found on base Orchestration operation")})
	}
	// Add project context
	if projectContextConf, found := orchestrateConfigurations.Get(libcompose.OPERATION_CONFIGURATION_LIBCOMPOSE_CONTEXT); found {
		if !projectContextConf.Set(handler.settings.Context) {
			result.Set(false, []error{errors.New("Could not set base libcompose net context.  Config set error on base Orchestration operation")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not set base libcompose net context.  Config not found on base Orchestration operation")})
	}
	// Add Stdout as an output writer
	if projectOutputConf, found := orchestrateConfigurations.Get(libcompose.OPERATION_CONFIGURATION_LIBCOMPOSE_OUTPUT); found {
		if !projectOutputConf.Set(os.Stdout) {
			result.Set(false, []error{errors.New("Could not set base libcompose output handler.  Config set error on base Orchestration operation")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not set base libcompose output handler.  Config not found on base Orchestration operation")})
	}
	if projectErrorConf, found := orchestrateConfigurations.Get(libcompose.OPERATION_CONFIGURATION_LIBCOMPOSE_ERROR); found {
		if !projectErrorConf.Set(os.Stderr) {
			result.Set(false, []error{errors.New("Could not set base libcompose error handler.  Config set error on base Orchestration operation")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not set base libcompose error handler.  Config not found on base Orchestration operation")})
	}

	return baseLibcomposeOrchestrate
}
