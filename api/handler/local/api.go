package local

import (
	// log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/james-nesbitt/wundertools-go/api"
	"github.com/james-nesbitt/wundertools-go/api/handler"
	"github.com/james-nesbitt/wundertools-go/api/handler/bytesource"
	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/config"
	"github.com/james-nesbitt/wundertools-go/api/operation/setting"
)

// Make a Local based API object, based on a project path
func MakeLocalAPI(settings LocalAPISettings) *LocalAPI {
	localAPI := LocalAPI{
		settings: &settings,
	}

	/**
	 * Build a local API by first building a config handler
	 * then a relates settings handler, and then using those
	 * to build the rest of the handlers
	 */

	// Create a base handler, which just wraps the settings
	local_base := LocalHandler_Base{
		settings:   &settings,
		operations: &operation.Operations{},
	}

	// Build a config wrapper using the base for settings
	local_config := LocalHandler_Config{
		LocalHandler_Base: local_base,
	}
	local_config.Init()
	localAPI.AddHandler(handler.Handler(&local_config))
	// Get a config wrapper for other handlers
	localAPI.Config = local_config.ConfigWrapper()

	// Build a settings wrapper which uses the configwrapper and the base
	local_settings := LocalHandler_Setting{
		LocalHandler_Base: local_base,
	}
	local_settings.SetConfigWrapper(localAPI.Config)
	local_settings.Init()
	localAPI.AddHandler(handler.Handler(&local_settings))
	// Get a settings wrapper for other handlers
	localAPI.Settings = local_settings.SettingWrapper()

	// Build an orchestration handler
	local_orchestration := LocalHandler_Orchestrate{
		LocalHandler_Base: local_base,
	}
	local_orchestration.SetSettingWrapper(localAPI.Settings)
	local_orchestration.Init()
	localAPI.AddHandler(handler.Handler(&local_orchestration))

	return &localAPI
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

	Config   config.ConfigWrapper
	Settings setting.SettingWrapper
}

// Validate the local API instance
func (api *LocalAPI) Validate() bool {
	return true
}
