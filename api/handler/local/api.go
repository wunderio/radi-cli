package local

import (
	log "github.com/Sirupsen/logrus"

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
		for _, err := range errs {
			log.WithError(err).Error(err.Error())
		}
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
}

// An API based entirely on local handler
type LocalAPI struct {
	api.BaseAPI
	settings *LocalAPISettings
}
