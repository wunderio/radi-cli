package operation

// BaseOperation a simple operation base class, which provides string methods via local variables
type BaseOperation struct {
	id          string
	label       string
	description string

	properties *Properties
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
func (operation *BaseOperation) Properties() *Properties {
	return operation.properties
}
