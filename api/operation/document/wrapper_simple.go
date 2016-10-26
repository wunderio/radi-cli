package document

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * A simple blocking wrapper for document operations
 */

// Constructor for SimpleDocumentWrapper
func New_SimpleDocumentWrapper(operations *operation.Operations) *SimpleDocumentWrapper {
	return &SimpleDocumentWrapper{
		operations: operations,
	}
}

// Simple blocking wrapper for document operations
type SimpleDocumentWrapper struct {
	operations *operation.Operations
}

func (wrapper *SimpleDocumentWrapper) Get(key string) (string, error) {
	return "", nil
}

func (wrapper *SimpleDocumentWrapper) Set(key string, doc string) error {
	return nil
}

func (wrapper *SimpleDocumentWrapper) List(parent string) ([]string, error) {
	return []string{}, nil
}
