package operation

/**
 * This file holds the definition for an API Operation, and
 * also defines a usefull Operation list struct, as well as
 * a utility BaseOperation struct, which can be used for
 * Operation inheritance.
 */

// A single operation
type Operation interface {

	// METADATA

	// Return the string machinename/id of the Operation
	Id() string
	// Return a user readable string label for the Operation
	Label() string
	// return a multiline string description for the Operation
	Description() string

	// Is this operation meant to be used only inside the API
	Internal() bool

	// FUNCTIONAL

	// Run a validation check on the Operation
	Validate() bool

	// What settings/values does the Operation provide to an implemenentor
	Properties() *Properties

	// Execute the Operation
	Exec() Result

	/**

	// OPERATIONAL

	Exec()

	// Operational status of the operation
	Status() (chan int)

	// Trigger for operation failure
	Fail() (chan bool)

	// Error list
	Errors() []error

	*/
}

// Operations are a keyed map of individual Operations
type Operations struct {
	opMap map[string]Operation
	order []string
}

// Add a new Operation to the map
func (operations *Operations) Add(add Operation) bool {
	if operations.opMap == nil {
		operations.opMap = map[string]Operation{}
		operations.order = []string{}
	}
	addId := add.Id()
	operations.opMap[addId] = add
	operations.order = append(operations.order, addId)
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
	if operations.opMap != nil {
		operation, ok := operations.opMap[id]
		return operation, ok
	} else {
		return nil, false
	}
}

// Order returns a slice of operation ids, used in iterators to maintain an operation order
func (operations *Operations) Order() []string {
	return operations.order
}
