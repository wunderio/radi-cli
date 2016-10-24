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
	BaseLibcomposeNameFilesOperation

	properties *operation.Properties
}

// Use a different Id() than the parent
func (logs *LibcomposeMonitorLogsOperation) Id() string {
	return OPERATION_ID_COMPOSE_MONITOR_LOGS
}

// Validate
func (logs *LibcomposeMonitorLogsOperation) Validate() bool {
	return true
}

// Provide static properties for the operation
func (logs *LibcomposeMonitorLogsOperation) Properties() *operation.Properties {
	if logs.properties == nil {
		newProperties := &operation.Properties{}
		newProperties.Merge(*logs.BaseLibcomposeStayAttachedOperation.Properties())
		newProperties.Merge(*logs.BaseLibcomposeNameFilesOperation.Properties())
		logs.properties = newProperties
	}
	return logs.properties
}

// Execute the libCompose monitor logs operation
func (logs *LibcomposeMonitorLogsOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	properties := logs.Properties()
	// pass all confs to make a project
	project, _ := MakeComposeProject(properties)

	// some confs we will use locally

	var netContext context.Context
	// net context
	if netContextProp, found := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_CONTEXT); found {
		netContext = netContextProp.Get().(context.Context)
	} else {
		result.Set(false, []error{errors.New("Libcompose up operation is missing the context property")})
	}

	var follow bool
	// follow conf
	if followProp, found := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_ATTACH_FOLLOW); found {
		follow = followProp.Get().(bool)
	} else {
		result.Set(true, []error{errors.New("Libcompose logs operation is missing the follow property")})
	}

	// output handling test
	if outputProp, found := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_OUTPUT); found {
		outputProp.Set(io.Writer(os.Stdout))
	}

	if success, _ := result.Success(); success {
		if err := project.APIProject.Log(netContext, follow); err != nil {
			result.Set(false, []error{err, errors.New("Could not attach to the project for logs")})
		}
	}

	return operation.Result(&result)
}
