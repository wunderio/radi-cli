package config

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
func (app *Application) AlterService(service *libCompose_config.ServiceConfig) {

	app.alterService_RewriteMappedVolumes(service)

}

// rewrite mapped service volumes to use app points.
func (app *Application) alterService_RewriteMappedVolumes(service *libCompose_config.ServiceConfig) {

	// short cut to the application paths, which we will use for substitution
	appPaths := app.Paths

	for index, volume := range service.Volumes {
		if !strings.Contains(volume, ":") {
			// this volume is not mapped
			continue
		}

		/**
		 * use the volume as a slice:
		 *   [0] : local path
		 *   [1] : container path
		 *   [2] : optional RO flag
		 */
		volumeSlices := strings.SplitN(volume, ":", 3)
		modified := false

		switch volumeSlices[0][0] {
		/**
		 * @TODO refactor this string comparison to be less cumbersome
		 */
		case []byte("~")[0]:
			homePath, _ := appPaths.Path("user-home")
			volumeSlices[0] = strings.Replace(volumeSlices[0], "~", homePath, 1)
			modified = true

		case []byte(".")[0]:
			appPath, _ := appPaths.Path("project-root")
			volumeSlices[0] = strings.Replace(volumeSlices[0], "~", appPath, 1)
			modified = true

		case []byte("@")[0]:
			alias := volumeSlices[0]
			if aliasPath, found := appPaths.Path(alias[1:]); found {
				volumeSlices[0] = strings.Replace(volumeSlices[0], alias, aliasPath, 1)
				modified = true
			}
		}

		if modified {
			service.Volumes[index] = strings.Join(volumeSlices, ":")
		}
	}

}
