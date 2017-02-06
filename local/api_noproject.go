package local

import (
	log "github.com/Sirupsen/logrus"

	api_api "github.com/wunderkraut/radi-api/api"
	api_builder "github.com/wunderkraut/radi-api/builder"
	handler_local "github.com/wunderkraut/radi-handlers/local"
	handler_null "github.com/wunderkraut/radi-handlers/null"
)

/**
 * If we have discovered that there is no local project folder,
 * then we will enable a minimum API, which can be used to create
 * a local folder
 *
 * The resulting API will be used to bootstrap the app.
 */

// Construct a LocalAPI with without expecting any local configuration
func MakeLocalAPI_NoProject(settings handler_local.LocalAPISettings) (api_api.API, error) {
	log.Debug("Local:API:: Building No-Project API")
	bootstrapApi := api_builder.BuilderAPI{}
	bootstrapApi.AddBuilder(handler_local.New_LocalBuilder(settings))

	// allow local project operations, which could be used to build a project
	bootstrapApi.ActivateBuilder("local", *api_builder.New_Implementations([]string{"project"}), nil)

	// Use null wrappers for those handlers that we don't have (to play safe)
	bootstrapApi.AddBuilder(handler_null.New_NullBuilder())
	bootstrapApi.ActivateBuilder("null", *api_builder.New_Implementations([]string{"config", "seting", "command"}), nil)

	return api_api.API(&bootstrapApi), nil
}
