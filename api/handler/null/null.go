package null

/**
 * The null implementation centers around the null handler, which
 * provides Operation implementations that do nothing, and are
 * generally safe to use in cases where an Operation is kind of
 * expected, but you don't have a handler to fit the role.
 */

import (
	"github.com/james-nesbitt/wundertools-go/api"
	"github.com/james-nesbitt/wundertools-go/api/handler"
)

// Api Constructor for which uses only a Null Handler
func MakeNullAPI() api.API {
	nullAPI := api.BaseAPI{}

	nullAPI.AddHandler(handler.Handler(NewNullHandler()))

	return api.API(&nullAPI)
}
