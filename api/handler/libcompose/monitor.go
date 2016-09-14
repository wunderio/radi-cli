package libcompose

import (
	"errors"
	"io"
	"os"

	"golang.org/x/net/context"

	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/monitor"
)

/**
 * Monitoring operations provided by the libcompose
 * handler.
 */

const (
	OPERATION_ID_COMPOSE_MONITOR_LOGS = monitor.OPERATION_ID_MONITOR_LOGS + ".compose"
)

// An operations which streams the container logs from libcompose
type LibcomposeMonitorLogsOperation struct {
	monitor.BaseMonitorLogsOperation
	BaseLibcomposeStayAttachedOperation
	BaseLibcomposeOrchestrateNameFilesOperation

	configurations *operation.Configurations
}

// Use a different Id() than the parent
func (logs *LibcomposeMonitorLogsOperation) Id() string {
	return OPERATION_ID_COMPOSE_MONITOR_LOGS
}

// Validate
func (logs *LibcomposeMonitorLogsOperation) Validate() bool {
	return true
}

// Provide static configurations for the operation
func (logs *LibcomposeMonitorLogsOperation) Configurations() *operation.Configurations {
	if logs.configurations == nil {
		newConfigurations := &operation.Configurations{}
		newConfigurations.Merge(*logs.BaseLibcomposeStayAttachedOperation.Configurations())
		newConfigurations.Merge(*logs.BaseLibcomposeOrchestrateNameFilesOperation.Configurations())
		logs.configurations = newConfigurations
	}
	return logs.configurations
}

// Execute the libCompose monitor logs operation
func (logs *LibcomposeMonitorLogsOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	configurations := logs.Configurations()
	// pass all confs to make a project
	project, _ := MakeComposeProject(configurations)

	// some confs we will use locally

	var netContext context.Context
	// net context
	if netContextConf, found := configurations.Get(OPERATION_CONFIGURATION_LIBCOMPOSE_CONTEXT); found {
		netContext = netContextConf.Get().(context.Context)
	} else {
		result.Set(false, []error{errors.New("Libcompose up operation is missing the context configuration")})
	}

	var follow bool
	// follow conf
	if followConf, found := configurations.Get(OPERATION_CONFIGURATION_LIBCOMPOSE_ATTACH_FOLLOW); found {
		follow = followConf.Get().(bool)
	} else {
		result.Set(true, []error{errors.New("Libcompose logs operation is missing the follow configuration")})
	}

	// output handling test
	if outputConf, found := configurations.Get(OPERATION_CONFIGURATION_LIBCOMPOSE_OUTPUT); found {
		outputConf.Set(io.Writer(os.Stdout))
	}

	if success, _ := result.Success(); success {
		if err := project.APIProject.Log(netContext, follow); err != nil {
			result.Set(false, []error{err, errors.New("Could not attach to the project for logs")})
		}
	}

	return operation.Result(&result)
}
