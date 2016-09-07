package document

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * List all topics or subtopics using the Documentation Handler
 */

// Base class for documentation topic list Operation
type BaseDocumentTopicListOperation struct{}

// Id the operation
func (list *BaseDocumentTopicListOperation) Id() string {
	return "document.list"
}

// Label the operation
func (list *BaseDocumentTopicListOperation) Label() string {
	return "Documentation topic list"
}

// Description for the operation
func (list *BaseDocumentTopicListOperation) Description() string {
	return "List document topics."
}

// Is this an internal API operation
func (list *BaseDocumentTopicListOperation) Internal() bool {
	return false
}
func (list *BaseDocumentTopicListOperation) Configurations() *operation.Configurations {
	return &operation.Configurations{}
}
