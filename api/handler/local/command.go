package local

import (
	"github.com/james-nesbitt/wundertools-go/api/handler/libcompose"
	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/command"
)

/**
 * Command operations for local projects
 */

// A handler for local command
type LocalHandler_Command struct {
	LocalHandler_Base
	LocalHandler_ConfigWrapperBase
	libcompose.BaseLibcomposeHandler
}

// Identify the handler
func (handler *LocalHandler_Command) Id() string {
	return "local.command"
}

// [Handler.]Init tells the LocalHandler_Orchestrate to prepare it's operations
func (handler *LocalHandler_Command) Init() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	ops := operation.Operations{}

	// Get shared base operation from the base handler
	baseLibcompose := *handler.BaseLibcomposeHandler.LibComposeBaseOp

	// Make a wrapper for the Command Config interpretation, based on itnerpreting YML settings
	wrapper := libcompose.CommandConfigWrapper(libcompose.New_BaseCommandConfigWrapperYmlOperation(handler.ConfigWrapper()))

	ops.Add(operation.Operation(&libcompose.LibcomposeCommandListOperation{BaseLibcomposeNameFilesOperation: baseLibcompose, Wrapper: wrapper}))
	ops.Add(operation.Operation(&libcompose.LibcomposeCommandGetOperation{BaseLibcomposeNameFilesOperation: baseLibcompose, Wrapper: wrapper}))

	handler.operations = &ops

	return operation.Result(&result)
}

// Make OrchestrateWrapper
func (handler *LocalHandler_Command) CommandWrapper() command.CommandWrapper {
	return command.New_SimpleCommandWrapper(handler.operations)
}
