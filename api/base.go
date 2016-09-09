package api

import (
	"github.com/james-nesbitt/wundertools-go/api/handler"
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * This file provides a simple API implementation called BaseAPI
 * which you can use in simple cases, where you want an API based
 * simply on collecting operations from handlers
 *
 * To use it, just make an instance, and use .AddHandler to add an
 * ordered queue of Handlers.  Operations will be collected from all
 * added handlers, later Handler Operations taking precendence over
 * prior.
 *
 * YOU DO NOT NEED TO USE THIS, IT IS OPTIONAL
 */

// BaseAPI is a base struct API implementation
type BaseAPI struct {
	handlers map[string]handler.Handler
}

// Validate returns true as along as at least one Handler has been added
func (base *BaseAPI) Validate() bool {
	return len(base.handlers) > 0
}

// AddHandler adds a Handler to the API, and will use it's Operations
func (base *BaseAPI) AddHandler(add handler.Handler) bool {
	if base.handlers == nil {
		base.handlers = map[string]handler.Handler{}
	}
	base.handlers[add.Id()] = add
	return true
}

// Handler retrieves a single keyed Handler from the list
func (base *BaseAPI) Handler(id string) (handler.Handler, bool) {
	handler, ok := base.handlers[id]
	return handler, ok
}

// Operations returns a list of all of the Operations provided by all of the Handlers
func (base *BaseAPI) Operations() operation.Operations {
	operations := operation.Operations{}
	for _, handler := range base.handlers {
		merge := handler.Operations()
		operations.Merge(merge)
	}
	return operations
}
