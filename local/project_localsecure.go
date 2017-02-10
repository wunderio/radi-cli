package local

import (
	log "github.com/Sirupsen/logrus"

	api_builder "github.com/wunderkraut/radi-api/builder"
	handler_local "github.com/wunderkraut/radi-handlers/local"
)

/**
 * Build a local SecureProject
 */

// Construct a LocalProject by checking some paths for the current user.
func MakeLocal_SecureProject(settings handler_local.LocalAPISettings) (api_builder.Project, error) {
	var err error

	// this is our actual local Project
	log.Debug("CLI:LocalProject: Building SecureProject")
	localProject := api_builder.New_SecureProject()

	return localProject.Project(), err

}
