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
	OPERATION_PROPERTY_LIBCOMPOSE_PROJECTNAME = "compose.projectname"
	// config for a project yml files
	OPERATION_PROPERTY_LIBCOMPOSE_COMPOSEFILES = "compose.composefiles"

	// Input/Output objects
	OPERATION_PROPERTY_LIBCOMPOSE_OUTPUT = "compose.output"
	OPERATION_PROPERTY_LIBCOMPOSE_ERROR  = "compose.error"

	/**
	 * General Properties for most operations
	 */

	// config for an orchestration context limiter
	OPERATION_PROPERTY_LIBCOMPOSE_CONTEXT = "compose.context"
	// Should a process stay attached and follow?
	OPERATION_PROPERTY_LIBCOMPOSE_ATTACH_FOLLOW = "compose.attach.follow"

	/**
	 * Operation specific contexts
	 */

	// config for up orchestration compose settings
	OPERATION_PROPERTY_LIBCOMPOSE_SETTINGS_UP = "compose.up"
	// config for down orchestration compose settings
	OPERATION_PROPERTY_LIBCOMPOSE_SETTINGS_DOWN = "compose.down"
)

/**
 * Properties which the libCompose handler uses
 */

// Project Name Property for a docker.libCompose project
type LibcomposeProjectnameProperty struct {
	operation.StringProperty
}

// Id for the Property
func (name *LibcomposeProjectnameProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_PROJECTNAME
}

// Label for the Property
func (name *LibcomposeProjectnameProperty) Label() string {
	return "Project name"
}

// Description for the Property
func (name *LibcomposeProjectnameProperty) Description() string {
	return "Compose project name, which is used in container, volume and network naming."
}

// Is the Property internal only
func (name *LibcomposeProjectnameProperty) Internal() bool {
	return false
}

// YAML file list Property for a docker.libCompose project
type LibcomposeComposefilesProperty struct {
	operation.StringSliceProperty
}

// Id for the Property
func (files *LibcomposeComposefilesProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_COMPOSEFILES
}

// Label for the Property
func (files *LibcomposeComposefilesProperty) Label() string {
	return "docker-compose yml file list"
}

// Description for the Property
func (files *LibcomposeComposefilesProperty) Description() string {
	return "An ordered list of docker-compose yml files, which are passed to libcompose."
}

// Is the Property internal only
func (files *LibcomposeComposefilesProperty) Internal() bool {
	return false
}

// A libcompose Property for net context limiting
type LibcomposeContextProperty struct {
	operation.ContextProperty
}

// Id for the Property
func (contextConf *LibcomposeContextProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_CONTEXT
}

// Label for the Property
func (contextConf *LibcomposeContextProperty) Label() string {
	return "context limiter"
}

// Description for the Property
func (contextConf *LibcomposeContextProperty) Description() string {
	return "A golang.org/x/net/context for controling execution."
}

// Is the Property internal only
func (contextConf *LibcomposeContextProperty) Internal() bool {
	return false
}

// Output handler Property for a docker.libCompose project
type LibcomposeOutputProperty struct {
	operation.WriterProperty
}

// Id for the Property
func (output *LibcomposeOutputProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_OUTPUT
}

// Label for the Property
func (output *LibcomposeOutputProperty) Label() string {
	return "Output writer"
}

// Description for the Property
func (output *LibcomposeOutputProperty) Description() string {
	return "Output io.Writer which will receive compose output from containers."
}

// Is the Property internal only
func (output *LibcomposeOutputProperty) Internal() bool {
	return false
}

// Error handler Property for a docker.libCompose project
type LibcomposeErrorProperty struct {
	operation.WriterProperty
}

// Id for the Property
func (err *LibcomposeErrorProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_ERROR
}

// Label for the Property
func (err *LibcomposeErrorProperty) Label() string {
	return "Error writer"
}

// Description for the Property
func (err *LibcomposeErrorProperty) Description() string {
	return "Error io.Writer which will receive compose output from containers."
}

// Is the Property internal only
func (err *LibcomposeErrorProperty) Internal() bool {
	return false
}

/**
 * These Properties are wrappers for the various libCompose options
 * structs in https://github.com/docker/libcompose/blob/master/project/options/types.go
 */

// Property for a docker.libCompose project to indicate that a process hsould stay attached and follow
type LibcomposeAttachFollowProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (follow *LibcomposeAttachFollowProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_ATTACH_FOLLOW
}

// Label for the Property
func (follow *LibcomposeAttachFollowProperty) Label() string {
	return "Follow"
}

// Description for the Property
func (follow *LibcomposeAttachFollowProperty) Description() string {
	return "When capturing output, stay attached and follow the output?"
}

// Is the Property internal only
func (follow *LibcomposeAttachFollowProperty) Internal() bool {
	return false
}

// A libcompose Property for net context limiting
type LibcomposeOptionsUpProperty struct {
	value libCompose_options.Up
}

// Id for the Property
func (optionsConf *LibcomposeOptionsUpProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_SETTINGS_UP
}

// Label for the Property
func (optionsConf *LibcomposeOptionsUpProperty) Label() string {
	return "Up operation options"
}

// Description for the Property
func (optionsConf *LibcomposeOptionsUpProperty) Description() string {
	return "Options to configure the Up.  See github.com/docker/libcompose/project/options for more information."
}

// Is the Property internal only
func (optionsConf *LibcomposeOptionsUpProperty) Internal() bool {
	return false
}

// Give an idea of what type of value the property consumes
func (optionsConf *LibcomposeOptionsUpProperty) Type() string {
	return "github.com/docker/libcompose/project/options.Up"
}

func (optionsConf *LibcomposeOptionsUpProperty) Get() interface{} {
	return interface{}(optionsConf.value)
}
func (optionsConf *LibcomposeOptionsUpProperty) Set(value interface{}) bool {
	if converted, ok := value.(libCompose_options.Up); ok {
		optionsConf.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected github.com/docker/libcompose/project/options.Up")
		return false
	}
}

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
