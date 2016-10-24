package local

import (
	"github.com/james-nesbitt/wundertools-go/api/handler/libcompose"
	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/orchestrate"
)

// A handler for local orchestration using libcompose
type LocalHandler_Orchestrate struct {
	LocalHandler_Base
	LocalHandler_SettingWrapperBase
	libcompose.BaseLibcomposeHandler
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

	// Use discovered/default settings to build a base operation struct, to be share across orchestration operations
	baseLibcompose := *handler.BaseLibcomposeHandler.LibComposeBaseOp

	// Now we can add orchestration operations that use that Base class
	ops.Add(operation.Operation(&libcompose.LibcomposeMonitorLogsOperation{BaseLibcomposeNameFilesOperation: baseLibcompose}))
	ops.Add(operation.Operation(&libcompose.LibcomposeOrchestrateUpOperation{BaseLibcomposeNameFilesOperation: baseLibcompose}))
	ops.Add(operation.Operation(&libcompose.LibcomposeOrchestrateDownOperation{BaseLibcomposeNameFilesOperation: baseLibcompose}))

	handler.operations = &ops

	return operation.Result(&result)
}

// Make OrchestrateWrapper
func (handler *LocalHandler_Orchestrate) OrchestrateWrapper() orchestrate.OrchestrateWrapper {
	return orchestrate.New_SimpleOrchestrateWrapper(handler.operations)
}
