package operation

/**
 * This file holds the definition for an API Operation, and
 * also defines a usefull Operation list struct, as well as
 * a utility BaseOperation struct, which can be used for
 * Operation inheritance.
 */

// A single operation
type Operation interface {
	// Run a validation check on the Operation
	Validate() bool

	// Is this operation meant to be used only inside the API
	Internal() bool

	// What settings does the Operation provide to an implemenentor
	Configurations() *Configurations

	// Return the string machinename/id of the Operation
	Id() string
	// Return a user readable string label for the Operation
	Label() string
	// return a multiline string description for the Operation
	Description() string

	// Execute the Operation
	Exec() Result
}

// Operations are a keyed map of individual Operations
type Operations struct {
	operationsMap   map[string]Operation
	operationsOrder []string
}

// Add a new Operation to the map
func (operations *Operations) Add(add Operation) bool {
	if operations.operationsMap == nil {
		operations.operationsMap = map[string]Operation{}
		operations.operationsOrder = []string{}
	}
	addId := add.Id()
	operations.operationsMap[addId] = add
	operations.operationsOrder = append(operations.operationsOrder, addId)
	return true
}

// Merge one Operations set into the current set
func (operations *Operations) Merge(merge *Operations) {
	for _, operation := range merge.Order() {
		mergeOperation, _ := merge.Get(operation)
		operations.Add(mergeOperation)
	}
}

// Operation accessor by id
func (operations *Operations) Get(id string) (Operation, bool) {
	if operations.operationsMap != nil {
		operation, ok := operations.operationsMap[id]
		return operation, ok
	} else {
		return nil, false
	}
}

// OperationOrder returns a slice of operation ids, used in iterators to maintain an operation order
func (operations *Operations) Order() []string {
	return operations.operationsOrder
}
