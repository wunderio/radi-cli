package main

import (
	"github.com/urfave/cli"

	api_api "github.com/james-nesbitt/kraut-api"
	cli_local "github.com/james-nesbitt/kraut-cli/local"
)

/**
 * This implementation shortcuts to making a LocalAPI,
 * and using it to add commands and operations to the
 * App.
 * It is a shortcut because in the future we want the
 * API to be able to be locally configured, but built
 * from non-local components.  This is a first
 * implementation built just to be fast and easy.
 *
 * This is a first implementation and is meant to be
 * replaced, once we can get a better "builder"
 * approach.
 *
 * @TODO write a better implementation.
 */

// Assume a LocalAPI, and add operations and commands to an app
func AppLocalCommands(app *cli.App) error {

	// Make a local API instance
	local, err := cli_local.MakeLocalAPI()

	// Add any operations from the api to the app
	AppApiOperations(app, api_api.API(local))

	// Add any commands from the api CommandWrapper to the app
	AppWrapperCommands(app, local.Command)

	return err
}
