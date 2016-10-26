package project

const (
	OPERATION_ID_PROJECT_CREATE = "project.create"
)

/**
 * Generate a new project
 */

// Generate a new project operation
type ProjectCreateOperation struct{}

// Id the operation
func (create *ProjectCreateOperation) Id() string {
	return OPERATION_ID_PROJECT_CREATE
}

// Label the operation
func (create *ProjectCreateOperation) Label() string {
	return "Create new project"
}

// Description for the operation
func (create *ProjectCreateOperation) Description() string {
	return "Create a new project from a templating source."
}

// Is this an internal API operation
func (create *ProjectCreateOperation) Internal() bool {
	return false
}
