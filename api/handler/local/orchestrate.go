package local

import (
	"errors"
	"os"
	"path"

	// "github.com/james-nesbitt/wundertools-go/api/handler/bytesource"
	"github.com/james-nesbitt/wundertools-go/api/handler/libcompose"
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

// A handler for local orchestration using libcompose
type LocalHandler_Orchestrate struct {
	LocalHandler_Base
	LocalHandler_SettingWrapperBase
}

// [Handler.]Id returns a string ID for the handler
func (handler *LocalHandler_Orchestrate) Id() string {
	return "local.orchestrate"
}

// [Handler.]Init tells the LocalHandler_Orchestrate to prepare it's operations
func (handler *LocalHandler_Orchestrate) Init() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	ops := operation.Operations{}

	// Set a project name
	projectName := "default"
	if settingsProjectName, err := handler.SettingWrapper().Get("Project"); err == nil {
		projectName = settingsProjectName
	} else {
		result.Set(false, []error{errors.New("Could not set base libCompose project name.  Config value not found in handler config")})
	}

	// Where to get docker-composer files
	dockerComposeFiles := []string{}
	// add the root composer file
	dockerComposeFiles = append(dockerComposeFiles, path.Join(handler.settings.ProjectRootPath, "docker-compose.yml"))

	// What net context to use
	runContext := handler.settings.Context

	// Output and Error writers
	outputWriter := os.Stdout
	errorWriter := os.Stderr

	// Use discovered/default settings to build a base operation struct, to be share across orchestration operations
	baseLibcomposeOrchestrate, constructResult := libcompose.New_BaseLibcomposeOrchestrateNameFilesOperation(
		projectName,
		dockerComposeFiles,
		runContext,
		outputWriter,
		errorWriter,
	)
	if success, constructErrors := constructResult.Success(); !success {
		result.Set(false, constructErrors)
	}

	// Now we can add orchestration operations that use that Base class
	ops.Add(operation.Operation(&libcompose.LibcomposeMonitorLogsOperation{BaseLibcomposeOrchestrateNameFilesOperation: baseLibcomposeOrchestrate}))
	ops.Add(operation.Operation(&libcompose.LibcomposeOrchestrateUpOperation{BaseLibcomposeOrchestrateNameFilesOperation: baseLibcomposeOrchestrate}))
	ops.Add(operation.Operation(&libcompose.LibcomposeOrchestrateDownOperation{BaseLibcomposeOrchestrateNameFilesOperation: baseLibcomposeOrchestrate}))

	handler.operations = &ops

	return operation.Result(&result)
}
