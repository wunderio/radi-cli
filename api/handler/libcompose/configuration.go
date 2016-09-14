package libcompose

import (
	log "github.com/Sirupsen/logrus"

	libCompose_options "github.com/docker/libcompose/project/options"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

const (
	/**
	 * Configs used in all operations to build the libCompose project
	 */

	// config for a project name
	OPERATION_CONFIGURATION_LIBCOMPOSE_PROJECTNAME = "compose.projectname"
	// config for a project yml files
	OPERATION_CONFIGURATION_LIBCOMPOSE_COMPOSEFILES = "compose.composefiles"

	// Input/Output objects
	OPERATION_CONFIGURATION_LIBCOMPOSE_OUTPUT = "compose.output"
	OPERATION_CONFIGURATION_LIBCOMPOSE_ERROR  = "compose.error"

	/**
	 * General configurations for most operations
	 */

	// config for an orchestration context limiter
	OPERATION_CONFIGURATION_LIBCOMPOSE_CONTEXT = "compose.context"
	// Should a process stay attached and follow?
	OPERATION_CONFIGURATION_LIBCOMPOSE_ATTACH_FOLLOW = "compose.attach.follow"

	/**
	 * Operation specific contexts
	 */

	// config for up orchestration compose settings
	OPERATION_CONFIGURATION_LIBCOMPOSE_SETTINGS_UP = "compose.up"
	// config for down orchestration compose settings
	OPERATION_CONFIGURATION_LIBCOMPOSE_SETTINGS_DOWN = "compose.down"
)

/**
 * Configurations which the libCompose handler uses
 */

// Project Name Configuration for a docker.libCompose project
type LibcomposeProjectnameConfiguration struct {
	operation.StringConfiguration
}

// Id for the configuration
func (name *LibcomposeProjectnameConfiguration) Id() string {
	return OPERATION_CONFIGURATION_LIBCOMPOSE_PROJECTNAME
}

// Label for the configuration
func (name *LibcomposeProjectnameConfiguration) Label() string {
	return "Project name"
}

// Description for the configuration
func (name *LibcomposeProjectnameConfiguration) Description() string {
	return "Compose project name, which is used in container, volume and network naming."
}

// YAML file list Configuration for a docker.libCompose project
type LibcomposeComposefilesConfiguration struct {
	operation.StringSliceConfiguration
}

// Id for the configuration
func (files *LibcomposeComposefilesConfiguration) Id() string {
	return OPERATION_CONFIGURATION_LIBCOMPOSE_COMPOSEFILES
}

// Label for the configuration
func (files *LibcomposeComposefilesConfiguration) Label() string {
	return "docker-compose yml file list"
}

// Description for the configuration
func (files *LibcomposeComposefilesConfiguration) Description() string {
	return "An ordered list of docker-compose yml files, which are passed to libcompose."
}

// A libcompose configuration for net context limiting
type LibcomposeContextConfiguration struct {
	operation.ContextConfiguration
}

// Id for the configuration
func (contextConf *LibcomposeContextConfiguration) Id() string {
	return OPERATION_CONFIGURATION_LIBCOMPOSE_CONTEXT
}

// Label for the configuration
func (contextConf *LibcomposeContextConfiguration) Label() string {
	return "context limiter"
}

// Description for the configuration
func (contextConf *LibcomposeContextConfiguration) Description() string {
	return "A golang.org/x/net/context for controling execution."
}

// Output handler Configuration for a docker.libCompose project
type LibcomposeOutputConfiguration struct {
	operation.WriterConfiguration
}

// Id for the configuration
func (output *LibcomposeOutputConfiguration) Id() string {
	return OPERATION_CONFIGURATION_LIBCOMPOSE_OUTPUT
}

// Label for the configuration
func (output *LibcomposeOutputConfiguration) Label() string {
	return "Output writer"
}

// Description for the configuration
func (output *LibcomposeOutputConfiguration) Description() string {
	return "Output io.Writer which will receive compose output from containers."
}

// Error handler Configuration for a docker.libCompose project
type LibcomposeErrorConfiguration struct {
	operation.WriterConfiguration
}

// Id for the configuration
func (err *LibcomposeErrorConfiguration) Id() string {
	return OPERATION_CONFIGURATION_LIBCOMPOSE_ERROR
}

// Label for the configuration
func (err *LibcomposeErrorConfiguration) Label() string {
	return "Error writer"
}

// Description for the configuration
func (err *LibcomposeErrorConfiguration) Description() string {
	return "Error io.Writer which will receive compose output from containers."
}

/**
 * These configurations are wrappers for the various libCompose options
 * structs in https://github.com/docker/libcompose/blob/master/project/options/types.go
 */

// Configuration for a docker.libCompose project to indicate that a process hsould stay attached and follow
type LibcomposeAttachFollowConfiguration struct {
	operation.BooleanConfiguration
}

// Id for the configuration
func (follow *LibcomposeAttachFollowConfiguration) Id() string {
	return OPERATION_CONFIGURATION_LIBCOMPOSE_ATTACH_FOLLOW
}

// Label for the configuration
func (follow *LibcomposeAttachFollowConfiguration) Label() string {
	return "Follow"
}

// Description for the configuration
func (follow *LibcomposeAttachFollowConfiguration) Description() string {
	return "When capturing output, stay attached and follow the output?"
}

// A libcompose configuration for net context limiting
type LibcomposeOptionsUpConfiguration struct {
	value libCompose_options.Up
}

// Id for the configuration
func (optionsConf *LibcomposeOptionsUpConfiguration) Id() string {
	return OPERATION_CONFIGURATION_LIBCOMPOSE_SETTINGS_UP
}

// Label for the configuration
func (optionsConf *LibcomposeOptionsUpConfiguration) Label() string {
	return "Up operation options"
}

// Description for the configuration
func (optionsConf *LibcomposeOptionsUpConfiguration) Description() string {
	return "Options to configure the Up.  See github.com/docker/libcompose/project/options for more information."
}

func (optionsConf *LibcomposeOptionsUpConfiguration) Get() interface{} {
	return interface{}(optionsConf.value)
}
func (optionsConf *LibcomposeOptionsUpConfiguration) Set(value interface{}) bool {
	if converted, ok := value.(libCompose_options.Up); ok {
		optionsConf.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Configuration value, because the passed parameter was the wrong type. Expected github.com/docker/libcompose/project/options.Up")
		return false
	}
}

// A libcompose configuration for net context limiting
type LibcomposeOptionsDownConfiguration struct {
	value libCompose_options.Down
}

// Id for the configuration
func (optionsConf *LibcomposeOptionsDownConfiguration) Id() string {
	return OPERATION_CONFIGURATION_LIBCOMPOSE_SETTINGS_DOWN
}

// Label for the configuration
func (optionsConf *LibcomposeOptionsDownConfiguration) Label() string {
	return "Down operation options"
}

// Description for the configuration
func (optionsConf *LibcomposeOptionsDownConfiguration) Description() string {
	return "Options to configure the Down.  See github.com/docker/libcompose/project/options for more information."
}

func (optionsConf *LibcomposeOptionsDownConfiguration) Get() interface{} {
	return interface{}(optionsConf.value)
}
func (optionsConf *LibcomposeOptionsDownConfiguration) Set(value interface{}) bool {
	if converted, ok := value.(libCompose_options.Down); ok {
		optionsConf.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Configuration value, because the passed parameter was the wrong type. Expected github.com/docker/libcompose/project/options.Down")
		return false
	}
}
