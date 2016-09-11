package handler

/**
 * This file defines the core handle functionality, which
 * actual handler implementations will implement.
 *
 */

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

type Handler interface {
	// Initialize and validate the Handler
	Init() operation.Result
	// Rturn a string identifier for the Handler (not functionally needed yet)
	Id() string
	// Return a list of Operations from the Handler
	Operations() *operation.Operations
}
