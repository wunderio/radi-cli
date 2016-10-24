package libcompose

/**
 * Alter a service, based on the app.
 *
 * This primarily offers a way to use short forms in
 * service definitions in yml files, but is is primarily
 * targets at operations not based on compose, as compose
 * has expected behaviour, and is piped right through
 * the libCompose code.
 */

import (
	"strings"

	libCompose_config "github.com/docker/libcompose/config"
)

/**
 * clean up a service based on this app
 */
func (project *ComposeProject) AlterService(service *libCompose_config.ServiceConfig) {
	project.alterService_RewriteMappedVolumes(service)
}

// rewrite mapped service volumes to use app points.
func (project *ComposeProject) alterService_RewriteMappedVolumes(service *libCompose_config.ServiceConfig) {

	for index, _ := range service.Volumes.Volumes {
		volume := service.Volumes.Volumes[index]

		switch volume.Source[0] {

		// relate volume to the current user home path
		case []byte("~")[0]:
			homePath := project.pathSettings.UserHomePath
			volume.Source = strings.Replace(volume.Source, "~", homePath, 1)

		// relate volume to project root
		case []byte(".")[0]:
			appPath := project.pathSettings.ProjectRootPath
			volume.Source = strings.Replace(volume.Source, "~", appPath, 1)

		// @TODO this is a stupid special hard-code that we should document somehow
		// @NOTE this is dangerous and will likely only work in cases where PWD is available
		case []byte("!")[0]:
			appPath := project.pathSettings.ExecPath
			volume.Source = strings.Replace(volume.Source, "!", appPath, 1)

		case []byte("@")[0]:
			if aliasPath, found := project.pathSettings.ConfigPaths.Get(volume.Source[1:]); found {
				volume.Source = strings.Replace(volume.Source, volume.Source, aliasPath.PathString(), 1)
			}
		}
	}

}
