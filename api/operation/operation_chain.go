package operation

// ChainOperation runs multiple operations in sequence.  Extend this and add ID/Label/Configurations handling
type ChainOperation struct {
	stopOnSuccess bool
	operations    *Operations
}

// Exec the chain operation by running Exec on each child
func (chain *ChainOperation) Exec() Result {
	chainResult := ChainResult{
		BaseResult{
			success: true,
			errors:  []error{},
		},
	}

	for _, id := range chain.operations.Order() {
		operation, _ := chain.operations.Get(id)
		result := operation.Exec()
		chainResult.AddResult(result)

		if resultSuccess, _ := result.Success(); chain.stopOnSuccess && resultSuccess {
			break
		}
	}

	return Result(&chainResult)
}
