package local

import (
	"errors"

	log "github.com/Sirupsen/logrus"

	api_api "github.com/wunderkraut/radi-api/api"
	api_builder "github.com/wunderkraut/radi-api/builder"
	api_config "github.com/wunderkraut/radi-api/operation/config"
	handlers_local "github.com/wunderkraut/radi-handlers/local"
)

/**
 * Make a local API
 *
 * First we check to see if we have a local project.  If not then
 * we return a minimal "build a project" api, which will give operations
 * but little real functionality outside of initializing a project
 *
 * If we have a project then we follow the following steps:
 *   1. build a bootstrap API, which is used to load configuraion
 *   2. Use the configuration from the boostrap to build a real
 *      API for the actual project config.
 *
 * This sequence is necessary to bypass the chicken-egg dilemna where
 * we need an API to get project settings, but we need project settings
 * to decide what API to build.  The Boostrap API does have some weight,
 * but not much to worry about (perhaps a couple of files are opened 2x)
 *
 */
func MakeLocalAPI(settings handlers_local.LocalAPISettings) (api_api.API, error) {

	if settings.ProjectDoesntExist {

		localAPI, _ := MakeLocalAPI_NoProject(settings)
		return localAPI, errors.New("No project found.")

	} else {

		// build an API with at least the config operations, which we will need for a config wrapper
		log.Debug("Local:API:: Building bootsrap API")
		bootstrapApi := api_builder.BuilderAPI{}
		bootstrapApi.AddBuilder(handlers_local.New_LocalBuilder(settings))
		bootstrapApi.ActivateBuilder("local", *api_builder.New_Implementations([]string{"config"}), nil)

		// Now use the config operations to determine what builds are needed
		bootstrapOps := bootstrapApi.Operations()
		bootstrapConfigWrapper := api_config.New_SimpleConfigWrapper(&bootstrapOps)

		localAPI, err := MakeLocalAPI_LocalSecure(settings, bootstrapConfigWrapper)
		// we could also use MakeLocalAPI_Local(settings, bootstrapConfigWrapper)

		return localAPI, err
	}

}
