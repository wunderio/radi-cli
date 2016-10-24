package libcompose

import (
	// libCompose_options "github.com/docker/libcompose/project/options"

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

	/**
	 * Operation specific contexts
	 */

	// config for up orchestration compose settings
	OPERATION_PROPERTY_LIBCOMPOSE_SETTINGS_UP = "compose.up"
	// config for down orchestration compose settings
	OPERATION_PROPERTY_LIBCOMPOSE_SETTINGS_DOWN = "compose.down"

	// Individual possible libcompose properties
	OPERATION_PROPERTY_LIBCOMPOSE_FORCEREMOVE      = "compose.forceremove"
	OPERATION_PROPERTY_LIBCOMPOSE_NOCACHE          = "compose.nocache"
	OPERATION_PROPERTY_LIBCOMPOSE_PULL             = "compose.pull"
	OPERATION_PROPERTY_LIBCOMPOSE_DETACH           = "compose.detach"
	OPERATION_PROPERTY_LIBCOMPOSE_NORECREATE       = "compose.norecreate"
	OPERATION_PROPERTY_LIBCOMPOSE_NOBUILD          = "compose.nobuild"
	OPERATION_PROPERTY_LIBCOMPOSE_FORCERECREATE    = "compose.forcerecreate"
	OPERATION_PROPERTY_LIBCOMPOSE_FORCEREBUILD     = "compose.forcerebuild"
	OPERATION_PROPERTY_LIBCOMPOSE_REMOVEVOLUMES    = "compose.removevolumes"
	OPERATION_PROPERTY_LIBCOMPOSE_REMOVEORPHANS    = "compose.removeorphans"
	OPERATION_PROPERTY_LIBCOMPOSE_REMOVEIMAGETYPES = "compose.removeimagetypes"
	OPERATION_PROPERTY_LIBCOMPOSE_REMOVERUNNING    = "compose.removerunning"
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

// BUILD : Property for a docker.libCompose project to indicate that a build should ignore cached image layers
type LibcomposeNoCacheProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (nocache *LibcomposeNoCacheProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_NOCACHE
}

// Label for the Property
func (nocache *LibcomposeNoCacheProperty) Label() string {
	return "nocache"
}

// Description for the Property
func (nocache *LibcomposeNoCacheProperty) Description() string {
	return "When capturing building, ignore cached docker layers?"
}

// Is the Property internal only
func (nocache *LibcomposeNoCacheProperty) Internal() bool {
	return false
}

// Property for a docker.libCompose project to indicate that a process remove ... something
type LibcomposeForceRemoveProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (forceremove *LibcomposeForceRemoveProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_FORCEREMOVE
}

// Label for the Property
func (forceremove *LibcomposeForceRemoveProperty) Label() string {
	return "Force remove"
}

// Description for the Property
func (forceremove *LibcomposeForceRemoveProperty) Description() string {
	return "When building, force remove .... something?"
}

// Is the Property internal only
func (forceremove *LibcomposeForceRemoveProperty) Internal() bool {
	return false
}

// Property for a docker.libCompose project to indicate that a process hsould stay attached and follow
type LibcomposePullProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (pull *LibcomposePullProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_PULL
}

// Label for the Property
func (pull *LibcomposePullProperty) Label() string {
	return "Pull"
}

// Description for the Property
func (pull *LibcomposePullProperty) Description() string {
	return "When building, pull all images before using them?"
}

// Is the Property internal only
func (pull *LibcomposePullProperty) Internal() bool {
	return false
}

// Property for a docker.libCompose project to indicate that a process hsould stay attached and follow
type LibcomposeDetachProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (detach *LibcomposeDetachProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_DETACH
}

// Label for the Property
func (detach *LibcomposeDetachProperty) Label() string {
	return "Detach"
}

// Description for the Property
func (detach *LibcomposeDetachProperty) Description() string {
	return "When capturing output, detach from the output?"
}

// Is the Property internal only
func (detach *LibcomposeDetachProperty) Internal() bool {
	return false
}

// UP : Property for a docker.libCompose project to indicate that a process should not create missing containers
type LibcomposeNoRecreateProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (norecreate *LibcomposeNoRecreateProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_NORECREATE
}

// Label for the Property
func (norecreate *LibcomposeNoRecreateProperty) Label() string {
	return "Create"
}

// Description for the Property
func (norecreate *LibcomposeNoRecreateProperty) Description() string {
	return "When starting a container, create it first, if it is missing?"
}

// Is the Property internal only
func (norecreate *LibcomposeNoRecreateProperty) Internal() bool {
	return false
}

// UP|RECREATE : Property for a docker.libCompose project to indicate that a process should build containers even if they are found
type LibcomposeForceRecreateProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (forcerecreate *LibcomposeForceRecreateProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_FORCERECREATE
}

// Label for the Property
func (forcerecreate *LibcomposeForceRecreateProperty) Label() string {
	return "Force Recreate"
}

// Description for the Property
func (forcerecreate *LibcomposeForceRecreateProperty) Description() string {
	return "Force recreating containers, even if they exist already?"
}

// Is the Property internal only
func (forcerecreate *LibcomposeForceRecreateProperty) Internal() bool {
	return false
}

// UP|CREATE : Property for a docker.libCompose project to indicate that a process should not build any containers
type LibcomposeNoBuildProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (dontbuild *LibcomposeNoBuildProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_NOBUILD
}

// Label for the Property
func (dontbuild *LibcomposeNoBuildProperty) Label() string {
	return "Don't Build"
}

// Description for the Property
func (dontbuild *LibcomposeNoBuildProperty) Description() string {
	return "Don't build any missing images?"
}

// Is the Property internal only
func (dontbuild *LibcomposeNoBuildProperty) Internal() bool {
	return false
}

// UP|CREATE : Property for a docker.libCompose project to indicate that a process should force rebuilding images
type LibcomposeForceRebuildProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (forcerebuild *LibcomposeForceRebuildProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_FORCEREBUILD
}

// Label for the Property
func (forcerebuild *LibcomposeForceRebuildProperty) Label() string {
	return "Force rebuild"
}

// Description for the Property
func (forcerebuild *LibcomposeForceRebuildProperty) Description() string {
	return "Force rebuilding any images, even if they exist already?"
}

// Is the Property internal only
func (forcerebuild *LibcomposeForceRebuildProperty) Internal() bool {
	return false
}

// DOWN|DELETE : Property for a docker.libCompose project to indicate that a process should remove any volumes
type LibcomposeRemoveVolumesProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (removevolumes *LibcomposeRemoveVolumesProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_REMOVEVOLUMES
}

// Label for the Property
func (removevolumes *LibcomposeRemoveVolumesProperty) Label() string {
	return "Remove volumes"
}

// Description for the Property
func (removevolumes *LibcomposeRemoveVolumesProperty) Description() string {
	return "When removing containers, remove any volumes?"
}

// Is the Property internal only
func (removevolumes *LibcomposeRemoveVolumesProperty) Internal() bool {
	return false
}

// DOWN : Property for a docker.libCompose project to indicate that a process should remove any orphan containers
type LibcomposeRemoveOrphansProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (removeorphans *LibcomposeRemoveOrphansProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_REMOVEORPHANS
}

// Label for the Property
func (removeorphans *LibcomposeRemoveOrphansProperty) Label() string {
	return "Remove orphans"
}

// Description for the Property
func (removeorphans *LibcomposeRemoveOrphansProperty) Description() string {
	return "When removing containers, remove any orphans?"
}

// Is the Property internal only
func (removeorphans *LibcomposeRemoveOrphansProperty) Internal() bool {
	return false
}

// DOWN : Property for a docker.libCompose project to indicate that a process should remove images of a certain type
type LibcomposeRemoveImageTypeProperty struct {
	operation.StringProperty
}

// Id for the Property
func (removeimagetypes *LibcomposeRemoveImageTypeProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_REMOVEIMAGETYPES
}

// Label for the Property
func (removeimagetypes *LibcomposeRemoveImageTypeProperty) Label() string {
	return "Remove image types"
}

// Description for the Property
func (removeimagetypes *LibcomposeRemoveImageTypeProperty) Description() string {
	return "When removing containers, remove either 'none' local' or 'all' images?"
}

// Is the Property internal only
func (removeimagetypes *LibcomposeRemoveImageTypeProperty) Internal() bool {
	return false
}

// DELETE : Property for a docker.libCompose project to indicate that a process should delete running containers
type LibcomposeRemoveRunningProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (removerunning *LibcomposeRemoveRunningProperty) Id() string {
	return OPERATION_PROPERTY_LIBCOMPOSE_REMOVERUNNING
}

// Label for the Property
func (removerunning *LibcomposeRemoveRunningProperty) Label() string {
	return "Remove running"
}

// Description for the Property
func (removerunning *LibcomposeRemoveRunningProperty) Description() string {
	return "When removing containers, remove running containers?"
}

// Is the Property internal only
func (removerunning *LibcomposeRemoveRunningProperty) Internal() bool {
	return false
}
