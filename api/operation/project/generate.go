package project

const (
	OPERATION_ID_PROJECT_GENERATE = "project.generate"
)

/**
 * Generate a new project
 */

// Generate a new project operation
type ProjectGenerateOperation struct{}

// Id the operation
func (generate *ProjectGenerateOperation) Id() string {
	return OPERATION_ID_PROJECT_GENERATE
}

// Label the operation
func (generate *ProjectGenerateOperation) Label() string {
	return "Generate new project template"
}

// Description for the operation
func (generate *ProjectGenerateOperation) Description() string {
	return "Genrate a new project template from the current project."
}

// Is this an internal API operation
func (generate *ProjectGenerateOperation) Internal() bool {
	return false
}
