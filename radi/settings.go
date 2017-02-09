package radi

import (
	"context"

	handler_bytesource "github.com/wunderkraut/radi-handlers/bytesource"
	handler_local "github.com/wunderkraut/radi-handlers/local"
)

// Construct a LocalAPI by checking some paths for the current user.
func MakeLocalAPISettings(workingDir string, ctx context.Context) handler_local.LocalAPISettings {
	// create an initial empty settings
	settings := handler_local.LocalAPISettings{
		BytesourceFileSettings: handler_bytesource.BytesourceFileSettings{
			ExecPath:    workingDir,
			ConfigPaths: &handler_bytesource.Paths{},
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
