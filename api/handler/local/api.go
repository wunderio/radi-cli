package local

import (
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/james-nesbitt/wundertools-go/api"
	"github.com/james-nesbitt/wundertools-go/api/handler"
	"github.com/james-nesbitt/wundertools-go/api/handler/bytesource"
)

// Make a Local based API object, based on a project path
func MakeLocalAPI(settings LocalAPISettings) api.API {
	localAPI := LocalAPI{
		settings: &settings,
	}

	localHandler := LocalHandler{
		settings: &settings,
	}
	if success, errs := localHandler.Init().Success(); !success {
		logger := log.WithFields(log.Fields{})
		for _, err := range errs {
			logger.Error(err.Error())
		}
		logger.Error("Failed to initialize local handler")
	}

	localAPI.AddHandler(handler.Handler(&localHandler))

	return api.API(&localAPI)
}

// Settings needed to make a local API
type LocalAPISettings struct {
	ProjectRootPath string
	UserHomePath    string
	ExecPath        string
	ConfigPaths     *bytesource.Paths
	Context         context.Context
}

// An API based entirely on local handler
type LocalAPI struct {
	api.BaseAPI
	settings *LocalAPISettings
}
