package project

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Properties used for various project operations
 */

const (
	// property id for create type
	OPERATION_PROPERTY_PROJECT_CREATE_TYPE = "project.create.type"
	// property id for create source
	OPERATION_PROPERTY_PROJECT_CREATE_SOURCE = "project.create.source"
)

// Property for the type of create
type ProjectCreateTypeProperty struct {
	operation.StringProperty
}

// Id for the Property
func (createType *ProjectCreateTypeProperty) Id() string {
	return OPERATION_PROPERTY_PROJECT_CREATE_TYPE
}

// Label for the Property
func (createType *ProjectCreateTypeProperty) Label() string {
	return "Create type."
}

// Description for the Property
func (createType *ProjectCreateTypeProperty) Description() string {
	return "Method used to create the project."
}

// Is the Property internal only
func (createType *ProjectCreateTypeProperty) Internal() bool {
	return false
}

// Property for the create source
type ProjectCreateSourceProperty struct {
	operation.StringProperty
}

// Id for the Property
func (createSource *ProjectCreateSourceProperty) Id() string {
	return OPERATION_PROPERTY_PROJECT_CREATE_SOURCE
}

// Label for the Property
func (createSource *ProjectCreateSourceProperty) Label() string {
	return "Create source."
}

// Description for the Property
func (createSource *ProjectCreateSourceProperty) Description() string {
	return "Method source used to create the project."
}

// Is the Property internal only
func (createSource *ProjectCreateSourceProperty) Internal() bool {
	return false
}
