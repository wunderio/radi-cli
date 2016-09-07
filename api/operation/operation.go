package operation 

/**
 * This file holds the definition for an API Operation, and
 * also defines a usefull Operation list struct, as well as
 * a utility BaseOperation struct, which can be used for 
 * Operation inheritance.
 */

// Operations are a keyed map of individual Operations
type Operations struct {
	operationsMap map[string]Operation
	operationsOrder []string
}

// Add a new Operation to the map
func (operations *Operations) Add(operation Operation) bool {
	operations.operationsMap[operation.Id()] = operation
	return true
}
// Merge one Operations set into the current set
func (operations *Operations) Merge(merge *Operations) {
	for _, operation := range merge.OperationOrder() {
		mergeOperation, _ := merge.Operation(operation)
		operations.Add(mergeOperation)
	}
}
// Operation accessor by id
func (operations *Operations) Operation(id string) (Operation, bool) {
	operation, ok := operations.operationsMap[id]
	return operation, ok
}
// OperationOrder returns a slice of operation ids, used in iterators to maintain an operation order
func (operations *Operations) OperationOrder() []string {
	return operations.operationsOrder
}

// A single operation
type Operation interface {
	// Run a validation check on the Operation
	Validate()    bool

	// Is this operation meant to be used only inside the API
	Internal()    bool

	// What settings does the Operation provide to an implemenentor
	Configurations() *Configurations

	// Return the string machinename/id of the Operation
	Id()          string
	// Return a user readable string label for the Operation
	Label()       string
	// return a multiline string description for the Operation
	Description() string

	// Execute the Operation
	Exec()        Result
}

// BaseOPeration a simple operation base class, which provides string methods via local variables
type BaseOperation struct {
	id             string
	label          string
	description    string

	configurations *Configurations
}

func (operation *BaseOperation) Validate() bool {
	return operation.Id() != ""
}
func (operation *BaseOperation) Id() string {
	return operation.id
}
func (operation *BaseOperation) Label() string {
	return operation.label
}
func (operation *BaseOperation) Description() string {
	return operation.description
}
func (operation *BaseOperation) Configurations() *Configurations {
	return operation.configurations
}

// ChainOperation runs multiple operations in sequence.  Extend this and add ID/Label/Configurations handling
type ChainOperation struct {
	stopOnSuccess  bool
	operations 	   *Operations
}
// Exec the chain operation by running Exec on each child
func (chain *ChainOperation) Exec() Result {
	chainResult := ChainResult{
		BaseResult{
			success: true,
			errors: []error{},
			},
	}

	for _, id := range chain.operations.OperationOrder() {
		operation, _ := chain.operations.Operation(id)
		result := operation.Exec()
		chainResult.AddResult(result)

		if resultSuccess, _ := result.Success(); chain.stopOnSuccess && resultSuccess {
			break
		}
	}

	return Result(&chainResult)
}
