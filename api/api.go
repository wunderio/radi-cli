package api

/**
 * The API is a stateful GO library, which can be used to connect
 * to a Wundertools APP concept, and provide a number of operations
 * that can be implemented on that app.
 *
 * The API is considered a static, but locally stateless library,
 * which can be created on the fly.  The API may manage state internally
 * but it's state doesn't need to be handled externally, except perhaps
 * by maintaining authentication tokens across uses (which we will
 * probably try to avoid.)
 *
 * The API works internally by use of various Handler implementations
 * which themselved define a number of keyed Operations.  The Operations
 * of the Handlers used by an API instance are collected and returned
 * on request, each of which is executable on it's own.
 *
 * Internally, the API Operation objects are abstract, but the keys
 * of certain Operations are of significance in some cases.  For
 * example, the authentication and user-retrieval operations are
 * used internally to enforce authorization control over other
 * operations, and the configuration retrieval may be used internally
 * in order to retrieve information about what other handlers should
 * be used.
 *
 */

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

// API is an instance of the API library that can return Operations
type API interface {
	// Validate returns a boolean value for if an API instance considers itself to be properly set up
	Validate() bool
	// Operations returns a list of executable operations
	Operations() operation.Operations
}
