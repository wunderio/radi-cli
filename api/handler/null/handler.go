package null

/**
 * The NullHandler provides a handlers with a set of operations that are
 * entirly Null provided.
 */

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/monitor"
)

// NullHandler Constructor, doesn't do much preprocessing really
func NewNullHandler() *NullHandler {
	nullHandler := NullHandler{}
	return &nullHandler
}

// NullHandler is a handler implementation that provides many core operations, but does very little (but is safe to use)
type NullHandler struct{}

// [Handler.]Id returns a string ID for the handler
func (handler *NullHandler) Id() string {
	return "null"
}

// [Handler.]Init tells the NullHandler to process itself. Return true as Null Handler always validates true
func (handler *NullHandler) Init() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)
	return operation.Result(&result)
}

// [Handler.]Operations returns an Operations list of a number of different Null operations
func (handler *NullHandler) Operations() *operation.Operations {
	operations := operation.Operations{}

	// Add Null config operations
	operations.Add(operation.Operation(&NullConfigReadersOperation{}))
	operations.Add(operation.Operation(&NullConfigWritersOperation{}))
	// Add Null setting operations
	operations.Add(operation.Operation(&NullSettingGetOperation{}))
	operations.Add(operation.Operation(&NullSettingSetOperation{}))
	// Add Null command operations
	operations.Add(operation.Operation(&NullCommandListOperation{}))
	operations.Add(operation.Operation(&NullCommandExecOperation{}))
	// Add Null documentation operations
	operations.Add(operation.Operation(&NullDocumentTopicListOperation{}))
	operations.Add(operation.Operation(&NullDocumentTopicGetOperation{}))
	// Add null monitor operations
	operations.Add(operation.Operation(&NullMonitorStatusOperation{}))
	operations.Add(operation.Operation(&NullMonitorInfoOperation{}))
	operations.Add(operation.Operation(&monitor.MonitorStandardLogOperation{}))
	// Add Null orchestration operations
	operations.Add(operation.Operation(&NullOrchestrateUpOperation{}))
	operations.Add(operation.Operation(&NullOrchestrateDownOperation{}))
	// Add Null security handlers
	operations.Add(operation.Operation(&NullSecurityAuthenticateOperation{}))
	operations.Add(operation.Operation(&NullSecurityAuthorizeOperation{}))
	operations.Add(operation.Operation(&NullSecurityUserOperation{}))

	return &operations
}
