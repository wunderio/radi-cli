package api 

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
func (operations *Operations) Operation(id string) (Operation, bool) {
	operation, ok := operations.operationsMap[id]
	return operation, ok
}
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
