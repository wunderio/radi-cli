package local

/**
 * Provide a Handler, and Operations, based on a local project, meaning
 * based on current local path, and the configuration files contained
 * therein (and perhaps also in the user home folder somewhere)
 */

import (
	"github.com/james-nesbitt/wundertools-go/api"
	"github.com/james-nesbitt/wundertools-go/api/handler"
)

// Api Constructor for which uses only a Local Handler
func MakeLocalAPI(PathProject string) api.API {
	local := api.BaseAPI{}

	local.AddHandler(handler.Handler(NewLocalHandler(PathProject)))

	return api.API(&local)
}
