package orchestrate

import (
	"errors"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * A simple, potentially blocking orchestrate wrapper implementation
 */

// Constructor for SimpleOrchestrateWrapper
func New_SimpleOrchestrateWrapper(operations *operation.Operations) *SimpleOrchestrateWrapper {
	return &SimpleOrchestrateWrapper{
		operations: operations,
	}
}

// A simple orchestration operation wrapper
type SimpleOrchestrateWrapper struct {
	operations *operation.Operations
}

// Orchestrate Up method
func (wrapper *SimpleOrchestrateWrapper) Up() error {
	var found, success bool
	var op operation.Operation
	var err []error

	if op, found = wrapper.operations.Get(OPERATION_ID_ORCHESTRATE_DOWN); !found {
		return errors.New("No up operation available in Orchestrate Wrapper")
	}

	if success, err = op.Exec().Success(); !success {
		return err[0] //errors.New("Operation get failed to execute in Setting Wrapper")
	}

	return nil
}

// Orchestrate Down method
func (wrapper *SimpleOrchestrateWrapper) Down() error {
	var found, success bool
	var op operation.Operation
	var err []error

	if op, found = wrapper.operations.Get(OPERATION_ID_ORCHESTRATE_DOWN); !found {
		return errors.New("No down operation available in Orchestrate Wrapper")
	}

	if success, err = op.Exec().Success(); !success {
		return err[0] //errors.New("Operation get failed to execute in Setting Wrapper")
	}

	return nil
}
