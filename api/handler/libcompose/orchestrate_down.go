package libcompose

import (
	"errors"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"

	libCompose_options "github.com/docker/libcompose/project/options"

	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/orchestrate"
)

/**
 * Down specific properties
 */

// A libcompose Property for net context limiting
type LibcomposeOptionsDownProperty struct {
	value libCompose_options.Down
}

// Id for the Property
func (optionsConf *LibcomposeOptionsDownProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_SETTINGS_DOWN
}

// Label for the Property
func (optionsConf *LibcomposeOptionsDownProperty) Label() string {
	return "Down operation options"
}

// Description for the Property
func (optionsConf *LibcomposeOptionsDownProperty) Description() string {
	return "Options to configure the Down.  See github.com/docker/libcompose/project/options for more information."
}

// Is the Property internal only
func (optionsConf *LibcomposeOptionsDownProperty) Internal() bool {
	return false
}

// Give an idea of what type of value the property consumes
func (optionsConf *LibcomposeOptionsDownProperty) Type() string {
	return "github.com/docker/libcompose/project/options.Down"
}

func (optionsConf *LibcomposeOptionsDownProperty) Get() interface{} {
	return interface{}(optionsConf.value)
}
func (optionsConf *LibcomposeOptionsDownProperty) Set(value interface{}) bool {
	if converted, ok := value.(libCompose_options.Down); ok {
		optionsConf.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected github.com/docker/libcompose/project/options.Down")
		return false
	}
}

/**
 * Operations
 */

// Base Down operation
type BaseLibcomposeOrchestrateDownSingleOperation struct {
	properties *operation.Properties
}

// Provide static Properties for the operation
func (base *BaseLibcomposeOrchestrateDownSingleOperation) Properties() *operation.Properties {
	if base.properties == nil {
		newProperties := &operation.Properties{}

		newProperties.Add(operation.Property(&LibcomposeOptionsDownProperty{}))

		base.properties = newProperties
	}
	return base.properties
}

// Base Down operation
type BaseLibcomposeOrchestrateDownParametrizedOperation struct {
	properties *operation.Properties
}

// Provide static Properties for the operation
func (base *BaseLibcomposeOrchestrateDownParametrizedOperation) Properties() *operation.Properties {
	if base.properties == nil {
		newProperties := &operation.Properties{}

		newProperties.Add(operation.Property(&LibcomposeRemoveVolumesProperty{}))
		newProperties.Add(operation.Property(&LibcomposeRemoveImageTypeProperty{}))
		newProperties.Add(operation.Property(&LibcomposeRemoveOrphansProperty{}))

		base.properties = newProperties
	}
	return base.properties
}

// LibCompose based down orchestrate operation
type LibcomposeOrchestrateDownOperation struct {
	orchestrate.BaseOrchestrationDownOperation
	BaseLibcomposeNameFilesOperation
	BaseLibcomposeOrchestrateDownParametrizedOperation

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
		down.properties.Merge(*down.BaseLibcomposeOrchestrateDownParametrizedOperation.Properties())
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
	var netContext context.Context
	var downOptions libCompose_options.Down

	// net context
	if netContextProp, found := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_CONTEXT); found {
		netContext = netContextProp.Get().(context.Context)
	} else {
		result.Set(false, []error{errors.New("Libcompose up operation is missing the context property")})
	}

	// up options
	downOptions = libCompose_options.Down{}
	if downOptionsProp, found := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_REMOVEVOLUMES); found {
		downOptions.RemoveVolume = downOptionsProp.Get().(bool)
	}
	if downOptionsProp, found := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_REMOVEIMAGETYPES); found {
		downOptions.RemoveImages = libCompose_options.ImageType(downOptionsProp.Get().(string))
	}
	if downOptionsProp, found := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_REMOVEORPHANS); found {
		downOptions.RemoveOrphans = downOptionsProp.Get().(bool)
	}

	if err := project.APIProject.Down(netContext, downOptions); err != nil {
		result.Set(false, []error{err})
	}

	return operation.Result(&result)
}
