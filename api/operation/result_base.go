package operation

// BaseResult is a base class for results which keep success boolean and errors slice as variables
type BaseResult struct {
	success bool
	errors  []error
}

// Set the state and add errors to the result
func (base *BaseResult) Set(success bool, errors []error) {
	if base.errors == nil {
		base.errors = []error{}
	}

	base.success = success
	if errors != nil {
		base.errors = append(base.errors, errors...)
	}
}
func (base *BaseResult) Success() (bool, []error) {
	return base.success, base.errors
}
