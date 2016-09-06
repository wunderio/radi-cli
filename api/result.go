package api

type Result interface {
	// Did the operation execute successfully? Return any error that occured
	Success()  (bool, []error)
}

type BaseResult struct {
	success bool
	errors  []error
}
func (result *BaseResult) Success() (bool, []error) {
	return result.success, result.errors
}
