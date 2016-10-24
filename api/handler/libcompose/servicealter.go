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
func (app *Application) AlterService(service *libCompose_config.ServiceConfig) {

	app.alterService_RewriteMappedVolumes(service)

}

// rewrite mapped service volumes to use app points.
func (app *Application) alterService_RewriteMappedVolumes(service *libCompose_config.ServiceConfig) {

	// short cut to the application paths, which we will use for substitution
	appPaths := app.Paths

	for index, _ := range service.Volumes.Volumes {

		volume := service.Volumes.Volumes[index]

		switch volume.Source[0] {
		/**
		 * @TODO refactor this string comparison to be less cumbersome
		 */
		case []byte("~")[0]:
			homePath, _ := appPaths.Path("user-home")
			volume.Source = strings.Replace(volume.Source, "~", homePath, 1)

		case []byte(".")[0]:
			appPath, _ := appPaths.Path("project-root")
			volume.Source = strings.Replace(volume.Source, "~", appPath, 1)

		case []byte("@")[0]:
			if aliasPath, found := appPaths.Path(volume.Source[1:]); found {
				volume.Source = strings.Replace(volume.Source, volume.Source, aliasPath, 1)
			}
		}
	}

}
