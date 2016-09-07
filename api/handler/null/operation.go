package null

/**
 * Operations that the Null Handler implements
 */

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/command"
	"github.com/james-nesbitt/wundertools-go/api/operation/config"
	"github.com/james-nesbitt/wundertools-go/api/operation/document"
	"github.com/james-nesbitt/wundertools-go/api/operation/monitor"
	"github.com/james-nesbitt/wundertools-go/api/operation/orchestrate"
	"github.com/james-nesbitt/wundertools-go/api/operation/security"
)

// Null base operation which always execs TRUE
type NullAllwaysTrueOperation struct{}

// Validate the operation
func (alwaystrue *NullAllwaysTrueOperation) Validate() bool {
	return true
}

// Exec the operation
func (alwaystrue *NullAllwaysTrueOperation) Exec() operation.Result {
	baseResult := operation.BaseResult{}
	baseResult.Set(true, []error{})
	return operation.Result(&baseResult)
}

/**
 * Command
 */

// Null operation for listing commands
type NullCommandListOperation struct {
	NullAllwaysTrueOperation
	command.BaseCommandListOperation
}

// Null operation for executing a command
type NullCommandExecOperation struct {
	NullAllwaysTrueOperation
	command.BaseCommandExecOperation
}

/**
 * Config
 */

// Null Configuration listing operation
type NullConfigListOperation struct {
	NullAllwaysTrueOperation
	config.BaseConfigListOperation
}

// Null Configuration retreive accessor operation
type NullConfigGetOperation struct {
	NullAllwaysTrueOperation
	config.BaseConfigGetOperation
}

// Null Configuration assign accessor operation
type NullConfigSetOperation struct {
	NullAllwaysTrueOperation
	config.BaseConfigSetOperation
}

/**
 * Documentationm
 */

// Null operation for listing documentation topics
type NullDocumentTopicListOperation struct {
	NullAllwaysTrueOperation
	document.BaseDocumentTopicListOperation
}

// Null Operation for retrieving a single documentation topic
type NullDocumentTopicGetOperation struct {
	NullAllwaysTrueOperation
	document.BaseDocumentTopicGetOperation
}

/**
 * Monitor
 */

// Null operation for monitoring information
type NullMonitorInfoOperation struct {
	NullAllwaysTrueOperation
	monitor.BaseMonitorInfoOperation
}

// Null status operation exec method
func (info *NullMonitorInfoOperation) Exec() operation.Result {
	message := "App is using NULL Info handler\n"
	info.WriteMessage(message)

	return info.NullAllwaysTrueOperation.Exec()
}

// Null operation for monitoring status
type NullMonitorStatusOperation struct {
	NullAllwaysTrueOperation
	monitor.BaseMonitorStatusOperation
}

// Null status operation exec method
func (status *NullMonitorStatusOperation) Exec() operation.Result {
	message := "App is using NULL status handler\n"
	status.WriteMessage(message)

	return status.NullAllwaysTrueOperation.Exec()
}

/**
 * Orchestration
 */

// Null operation for orchestration UP
type NullOrchestrateUpOperation struct {
	NullAllwaysTrueOperation
	orchestrate.BaseOrchestrationUpOperation
}

// Null operation for orchestration DOWN
type NullOrchestrateDownOperation struct {
	NullAllwaysTrueOperation
	orchestrate.BaseOrchestrationDownOperation
}

/**
 * Security
 */

// Null Authenticate always authenticates
type NullSecurityAuthenticateOperation struct {
	NullAllwaysTrueOperation
	security.BaseSecurityAuthenticateOperation
}

// Null Authorize always authorizes
type NullSecurityAuthorizeOperation struct {
	NullAllwaysTrueOperation
	security.BaseSecurityAuthorizeOperation
}

// Null User, provides a consistent user value
type NullSecurityUserOperation struct {
	NullAllwaysTrueOperation
	security.BaseSecurityUserOperation
}
