package document

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Retrieve the documentation for a single Documentation topic,
 * using the Documentation handler
 */

// Base class for documentation topic get Operation
type BaseDocumentTopicGetOperation struct{}

// Id the operation
func (get *BaseDocumentTopicGetOperation) Id() string {
	return "document.get"
}

// Label the operation
func (get *BaseDocumentTopicGetOperation) Label() string {
	return "Documentation topic get"
}

// Description for the operation
func (get *BaseDocumentTopicGetOperation) Description() string {
	return "List document topics."
}

// Is this an internal API operation
func (get *BaseDocumentTopicGetOperation) Internal() bool {
	return false
}
func (get *BaseDocumentTopicGetOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}
