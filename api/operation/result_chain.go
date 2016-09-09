package operation

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
