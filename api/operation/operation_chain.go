package operation

// ChainOperation runs multiple operations in sequence.  Extend this and add ID/Label/Configurations handling
type ChainOperation struct {
	// Tell this operation to stop processing the chained operations on the first TRUE result
	stopOnSuccess bool
	// The ordered chain of Operations to process when this operation is called
	operations *Operations
}

// Get Operation Configuration from all operations
func (chain *ChainOperation) Configurations() *Configurations {
	configurations := Configurations{}
	for _, key := range chain.operations.Order() {
		op, _ := chain.operations.Get(key)
		configurations.Merge(*op.Configurations())
	}
	return &configurations
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
