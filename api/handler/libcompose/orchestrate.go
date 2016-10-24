package libcompose

import (
	"errors"

	"golang.org/x/net/context"

	libCompose_options "github.com/docker/libcompose/project/options"

	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/orchestrate"
)

// LibCompose based up orchestrate operation
type LibcomposeOrchestrateUpOperation struct {
	orchestrate.BaseOrchestrationUpOperation
	BaseLibcomposeOrchestrateUpOperation
	BaseLibcomposeNameFilesOperation

	properties *operation.Properties
}

// Validate the libCompose Orchestrate Up operation
func (up *LibcomposeOrchestrateUpOperation) Validate() bool {
	return true
}

// Provide static properties for the operation
func (up *LibcomposeOrchestrateUpOperation) Properties() *operation.Properties {
	if up.properties == nil {
		newProperties := &operation.Properties{}
		newProperties.Merge(*up.BaseLibcomposeOrchestrateUpOperation.Properties())
		newProperties.Merge(*up.BaseLibcomposeNameFilesOperation.Properties())
		up.properties = newProperties
	}
	return up.properties
}

// Execute the libCompose Orchestrate Up operation
func (up *LibcomposeOrchestrateUpOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	properties := up.Properties()
	// pass all props to make a project
	project, _ := MakeComposeProject(properties)

	// some props we will use locally

	var netContext context.Context
	var upOptions libCompose_options.Up
	// net context
	if netContextProp, found := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_CONTEXT); found {
		netContext = netContextProp.Get().(context.Context)
	} else {
		result.Set(false, []error{errors.New("Libcompose up operation is missing the context property")})
	}

	// up options
	if upOptionsProp, found := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_SETTINGS_UP); found {
		upOptions = upOptionsProp.Get().(libCompose_options.Up)
	} else {
		result.Set(false, []error{errors.New("Libcompose up operation is missing the UP property")})
	}

	if success, _ := result.Success(); success {
		if err := project.APIProject.Up(netContext, upOptions); err != nil {
			result.Set(false, []error{err})
		}
	}

	return operation.Result(&result)
}

// LibCompose based down orchestrate operation
type LibcomposeOrchestrateDownOperation struct {
	orchestrate.BaseOrchestrationDownOperation
	BaseLibcomposeOrchestrateDownOperation
	BaseLibcomposeNameFilesOperation

	properties *operation.Properties
}

// Validate the libCompose Orchestrate Down operation
func (down *LibcomposeOrchestrateDownOperation) Validate() bool {
	return true
}

// Provide static properties for the operation
func (down *LibcomposeOrchestrateDownOperation) Properties() *operation.Properties {
	if down.properties == nil {
		down.properties = &operation.Properties{}
		down.properties.Merge(*down.BaseLibcomposeOrchestrateDownOperation.Properties())
		down.properties.Merge(*down.BaseLibcomposeNameFilesOperation.Properties())
	}
	return down.properties
}

// Execute the libCompose Orchestrate Down operation
func (down *LibcomposeOrchestrateDownOperation) Exec() operation.Result {
	result := operation.BaseResult{}

	properties := down.Properties()
	// pass all props to make a project
	project, _ := MakeComposeProject(properties)

	// some props we will use locally

	// net context
	netContextProp, _ := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_CONTEXT)
	netContext := netContextProp.Get().(context.Context)
	// down options
	downOptionsProp, _ := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_SETTINGS_DOWN)
	downOptions := downOptionsProp.Get().(libCompose_options.Down)

	if err := project.APIProject.Down(netContext, downOptions); err != nil {
		result.Set(false, []error{err})
	}

	return operation.Result(&result)
}
