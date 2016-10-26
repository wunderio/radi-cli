package operation

/**
 * Properties are abstract Operations values which
 * are meant to be accessed from outside of an operation
 * to either retrieve or assign.
 * This file provides a property collection struct,
 * and the interface for a single Property
 *
 * A Peroperty consumer should either recognize the
 * Operation by it's keys, and then handle it's Properties
 * as "knowns", or it should iterate through the Properties
 * and use some user-interface to allow interaction.
 *
 * A Property is typically used, by overwriting it's
 * value before running the operation, or by retreiving it's
 * value after the operation has executed.
 */

// A single Property for an operation
type Property interface {
	// ID returns string unique property Identifier
	Id() string
	// Label returns a short user readable label for the property
	Label() string
	// Description provides a longer multi-line string description of what the property does
	Description() string
	// Mark a property as being for internal use only (no shown to users)
	Internal() bool

	// Give an idea of what type of value the property consumes
	Type() string

	// Value allows the retrieval and setting of unknown Typed values for the property.
	Get() interface{}
	Set(interface{}) bool
}

// A set of Properties
type Properties struct {
	propMap map[string]Property
	order   []string
}

// safe initialization of vars
func (properties *Properties) makeSafe() {
	if properties.propMap == nil {
		properties.propMap = map[string]Property{}
		properties.order = []string{}
	}
}

// Add a property
func (properties *Properties) Add(property Property) {
	properties.makeSafe()
	id := property.Id()
	if _, exists := properties.propMap[id]; !exists {
		properties.order = append(properties.order, id)
	}
	properties.propMap[id] = property
}

// Merge in one set of properties into this configurations
func (properties *Properties) Merge(merge Properties) {
	for _, id := range merge.Order() {
		property, _ := merge.Get(id)
		properties.Add(property)
	}
}

// Retrieve a single property based on key id
func (properties *Properties) Get(id string) (Property, bool) {
	properties.makeSafe()
	property, ok := properties.propMap[id]
	return property, ok
}

// Retrieve and ordered list of property keys
func (properties *Properties) Order() []string {
	properties.makeSafe()
	return properties.order
}
