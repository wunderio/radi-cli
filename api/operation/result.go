package operation

/**
 * This file holds the definition of Result, which is what an operation
 * returns, and a few usefull base structs that implement Result, which
 * can be used directly or for inheritance
 */

// Result is an what an operation returns
type Result interface {
	// Did the operation execute successfully? Return any error that occured
	Success() (bool, []error)
}

// BaseResult is a base class for results which keep success boolean and errors slice as variables
type BaseResult struct {
	success bool
	errors  []error
}

// Set the state and add errors to the result
func (base *BaseResult) Set(success bool, errors []error) {
	base.success = success
	base.errors = append(base.errors, errors...)
}
func (base *BaseResult) Success() (bool, []error) {
	return base.success, base.errors
}

// ChainResult is a Result that aggregates multiple results
type ChainResult struct {
	BaseResult
}

// Add A result to the chain
func (chain *ChainResult) AddResult(add Result) {
	chainSuccess, chainErrors := chain.Success()
	addSuccess, addErrors := add.Success()

	chain.success = chainSuccess && addSuccess
	chain.errors = append(chainErrors, addErrors...)
}
