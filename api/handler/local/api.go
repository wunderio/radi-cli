package local

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"path"

	"github.com/james-nesbitt/wundertools-go/api"
	"github.com/james-nesbitt/wundertools-go/api/handler"
	"github.com/james-nesbitt/wundertools-go/api/handler/bytesource"
	"github.com/james-nesbitt/wundertools-go/api/handler/libcompose"
	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/command"
	"github.com/james-nesbitt/wundertools-go/api/operation/config"
	"github.com/james-nesbitt/wundertools-go/api/operation/orchestrate"
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

	/**
	 * We now have enough configuration/settings source
	 * where we can try to build our project, pulling
	 * in settings from these sources
	 */

	// Set a project name
	projectName := "default"
	if settingsProjectName, err := localAPI.Settings.Get("Project"); err == nil {
		projectName = settingsProjectName
	} else {
		log.WithError(errors.New("Could not set base libCompose project name.  Config value not found in handler config")).Error()
	}

	// Where to get docker-composer files
	dockerComposeFiles := []string{}
	// add the root composer file
	dockerComposeFiles = append(dockerComposeFiles, path.Join(settings.ProjectRootPath, "docker-compose.yml"))

	// What net context to use
	runContext := settings.Context

	// Output and Error writers
	outputWriter := os.Stdout
	errorWriter := os.Stderr

	/**
	 * Build a Handler base that produces LibCompose projects
	 * across other handlers which will then allow libCompose
	 * projects to be created
	 */

	// LibComposeHandlerBase
	base_libcompose := libcompose.New_BaseLibcomposeHandler(projectName, dockerComposeFiles, runContext, outputWriter, errorWriter, settings.BytesourceFileSettings)

	/**
	 * Build Handlers that actually provide usefull
	 * Operations now.
	 */

	// Build an orchestration handler
	local_orchestration := LocalHandler_Orchestrate{
		LocalHandler_Base:     local_base,
		BaseLibcomposeHandler: *base_libcompose,
	}
	local_orchestration.SetSettingWrapper(localAPI.Settings)
	local_orchestration.Init()
	localAPI.AddHandler(handler.Handler(&local_orchestration))
	// Get an orchestrate wrapper for other handlers
	localAPI.Orchestrate = local_orchestration.OrchestrateWrapper()

	// Build a command Handler
	local_command := LocalHandler_Command{
		LocalHandler_Base:     local_base,
		BaseLibcomposeHandler: *base_libcompose,
	}
	local_command.SetConfigWrapper(localAPI.Config)
	local_command.Init()
	localAPI.AddHandler(handler.Handler(&local_command))
	// Get an orchestrate wrapper for other handlers
	localAPI.Command = local_command.CommandWrapper()

	return &localAPI
}

// Settings needed to make a local API
type LocalAPISettings struct {
	bytesource.BytesourceFileSettings
	Context context.Context
}

// An API based entirely on local handler
type LocalAPI struct {
	api.BaseAPI
	settings *LocalAPISettings

	Command     command.CommandWrapper
	Config      config.ConfigWrapper
	Settings    setting.SettingWrapper
	Orchestrate orchestrate.OrchestrateWrapper
}

// Validate the local API instance
func (api *LocalAPI) Validate() bool {
	return true
}
