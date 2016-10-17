package operation

// A wrapping operation that decorates one operation with another
type DecoratedOperation struct {
	// Operation that wraps or decorates the other operation
	decorating Operation
	// Operation being decorated
	decorated Operation
}

// Get decorted operation id
func (operation *DecoratedOperation) Id() string {
	return operation.decorated.Id()
}

// Get decorted operation label
func (operation *DecoratedOperation) Label() string {
	return operation.decorated.Label() + " [" + operation.decorating.Label() + "]"
}

// Get Operation Properties from both operations
func (operation *DecoratedOperation) Properties() *Properties {
	Properties := operation.decorated.Properties()
	Properties.Merge(*operation.decorating.Properties())
	return Properties
}

// Execute the decorating operation, and then execute the decorated operation if the decorating was successful
func (operation *DecoratedOperation) Exec() Result {
	result := operation.decorating.Exec()
	if success, _ := result.Success(); !success {
		return result
	}
	return operation.decorated.Exec()
}
