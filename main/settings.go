package main

import (
	"golang.org/x/net/context"

	handlers_bytesource "github.com/wunderkraut/radi-handlers/bytesource"
	handlers_local "github.com/wunderkraut/radi-handlers/local"
)

// Construct a LocalAPI by checking some paths for the current user.
func MakeLocalAPISettings(workingDir string, ctx context.Context) handlers_local.LocalAPISettings {
	// create an initial empty settings
	settings := handlers_local.LocalAPISettings{
		BytesourceFileSettings: handlers_bytesource.BytesourceFileSettings{
			ExecPath:    workingDir,
			ConfigPaths: &handlers_bytesource.Paths{},
		},
		Context: ctx,
	}

	// }

	/**
	 * We could here add more paths for settings.ConfigPaths, for
	 * configurations of a higher priority.
	 */

	return settings
}
