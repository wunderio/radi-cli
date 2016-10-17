package operation

/**
 * A wrapper for a value, where the value could come from,
 * or go to different scopes, so the value is aware of it's
 * own scope
 */

// Scope value
type ScopedValue struct {
	val   []byte
	scope string
}

// Set value and scope
func (scoped *ScopedValue) Set(scope string, value []byte) {
	scoped.val = value
	scoped.scope = scope
}

// Get the value
func (scoped *ScopedValue) Value() []byte {
	return scoped.val
}

// Get the value as a string
func (scoped *ScopedValue) ValueString() string {
	return string(scoped.val)
}

// Get the value scope
func (scoped *ScopedValue) Scope() string {
	return scoped.scope
}
