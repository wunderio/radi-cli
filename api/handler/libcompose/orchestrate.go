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
	BaseLibcomposeOrchestrateNameFilesOperation

	configurations *operation.Configurations
}

// Validate the libCompose Orchestrate Up operation
func (up *LibcomposeOrchestrateUpOperation) Validate() bool {
	return true
}

// Provide static configurations for the operation
func (up *LibcomposeOrchestrateUpOperation) Configurations() *operation.Configurations {
	if up.configurations == nil {
		newConfigurations := &operation.Configurations{}
		newConfigurations.Merge(*up.BaseLibcomposeOrchestrateUpOperation.Configurations())
		newConfigurations.Merge(*up.BaseLibcomposeOrchestrateNameFilesOperation.Configurations())
		up.configurations = newConfigurations
	}
	return up.configurations
}

// Execute the libCompose Orchestrate Up operation
func (up *LibcomposeOrchestrateUpOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	configurations := up.Configurations()
	// pass all confs to make a project
	project, _ := MakeComposeProject(configurations)

	// some confs we will use locally

	var netContext context.Context
	var upOptions libCompose_options.Up
	// net context
	if netContextConf, found := configurations.Get(OPERATION_CONFIGURATION_LIBCOMPOSE_CONTEXT); found {
		netContext = netContextConf.Get().(context.Context)
	} else {
		result.Set(false, []error{errors.New("Libcompose up operation is missing the context configuration")})
	}

	// up options
	if upOptionsConf, found := configurations.Get(OPERATION_CONFIGURATION_LIBCOMPOSE_SETTINGS_UP); found {
		upOptions = upOptionsConf.Get().(libCompose_options.Up)
	} else {
		result.Set(false, []error{errors.New("Libcompose up operation is missing the UP configuration")})
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
	BaseLibcomposeOrchestrateNameFilesOperation

	configurations *operation.Configurations
}

// Validate the libCompose Orchestrate Down operation
func (down *LibcomposeOrchestrateDownOperation) Validate() bool {
	return true
}

// Provide static configurations for the operation
func (down *LibcomposeOrchestrateDownOperation) Configurations() *operation.Configurations {
	if down.configurations == nil {
		down.configurations = &operation.Configurations{}
		down.configurations.Merge(*down.BaseLibcomposeOrchestrateDownOperation.Configurations())
		down.configurations.Merge(*down.BaseLibcomposeOrchestrateNameFilesOperation.Configurations())
	}
	return down.configurations
}

// Execute the libCompose Orchestrate Down operation
func (down *LibcomposeOrchestrateDownOperation) Exec() operation.Result {
	result := operation.BaseResult{}

	configurations := down.Configurations()
	// pass all confs to make a project
	project, _ := MakeComposeProject(configurations)

	// some confs we will use locally

	// net context
	netContextConf, _ := configurations.Get(OPERATION_CONFIGURATION_LIBCOMPOSE_CONTEXT)
	netContext := netContextConf.Get().(context.Context)
	// down options
	downOptionsConf, _ := configurations.Get(OPERATION_CONFIGURATION_LIBCOMPOSE_SETTINGS_DOWN)
	downOptions := downOptionsConf.Get().(libCompose_options.Down)

	if err := project.APIProject.Down(netContext, downOptions); err != nil {
		result.Set(false, []error{err})
	}

	return operation.Result(&result)
}
