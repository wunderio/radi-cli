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
	Init()
	Validate() bool

	Id() string

	Operations() *operation.Operations	
}
