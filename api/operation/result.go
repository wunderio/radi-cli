package operation

/**
 * This file holds the definition of Result, which is what an operation
 * returns, and a few usefull base structs that implement Result, which
 * can be used directly or for inheritance
 */

// Result is an what an operation returns
type Result interface {
	// Did the operation execute successfully? Return any error that occured
	Success() (bool, []error)
}
