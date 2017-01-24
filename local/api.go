package local

import (
	"errors"
	"os"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"

	api_api "github.com/james-nesbitt/radi-api/api"
	api_builder "github.com/james-nesbitt/radi-api/builder"
	api_config "github.com/james-nesbitt/radi-api/operation/config"
	handlers_bytesource "github.com/james-nesbitt/radi-handlers/bytesource"
	handlers_configwrapper "github.com/james-nesbitt/radi-handlers/configwrapper"
	handlers_local "github.com/james-nesbitt/radi-handlers/local"
	handlers_null "github.com/james-nesbitt/radi-handlers/null"
	handlers_upcloud "github.com/james-nesbitt/radi-handlers/upcloud"
)

const (
	WUNDERTOOLS_PROJECT_CONF_FOLDER = ".radi" // If the project has existing setitngs, they will be in this subfolder, somewhere up the file tree.
	WUNDERTOOLS_USER_CONF_SUBPATH   = "radi"  // If the user has user-scope config, they will be in this subfolder
)

/**
 * Build a local API, by scanning for project settings based on the
 * path.  First a number of "conf" folders are determinged, and these
 * are used to build the localAPI.
 */

// Construct a LocalAPI by checking some paths for the current user.
func MakeLocalAPI() (api_api.API, error) {
	var err error

	workingDir, _ := os.Getwd()
	settings := handlers_local.LocalAPISettings{
		BytesourceFileSettings: handlers_bytesource.BytesourceFileSettings{
			ExecPath:    workingDir,
			ConfigPaths: &handlers_bytesource.Paths{},
		},
		Context: context.Background(),
	}

	// Discover paths for the user like ~ and ~/.config/wundertools
	DiscoverUserPaths(&settings)
	DiscoverProjectPaths(&settings)

	/**
	 * We could here add more paths for settings.ConfigPaths, for
	 * configurations of a higher priority.  For example, a feature
	 * or environment concept might want to override user and
	 * project level confs
	 */

	/**
	 * Now that we have local settings, let's start to build our API
	 *
	 * We will build it using a BuilderAPI, and adding the local
	 * Builder, which will be used at a minimum for a ConfigWrapper,
	 * to determine how to build the rest of the project.
	 * This API will be used to bootstrap the app.
	 */
	log.Debug("CLI:LocalAPI: Building bootsrap API")
	bootstrapApi := api_builder.BuilderAPI{}
	bootstrapApi.AddBuilder(handlers_local.New_LocalBuilder(settings))

	/**
	 * If we have discovered that there is no local project folder,
	 * then we will enable a minimum API, which can be used to create
	 * a local folder
	 */
	if settings.ProjectDoesntExist {

		// allow local project operations, which could be used to build a project
		bootstrapApi.ActivateBuilder("local", *api_builder.New_Implementations([]string{"project"}), nil)

		// Use null wrappers for those handlers that we don't have (to play safe)
		bootstrapApi.AddBuilder(handlers_null.New_NullBuilder())
		bootstrapApi.ActivateBuilder("null", *api_builder.New_Implementations([]string{"config", "seting", "command"}), nil)

		err = errors.New("No project found.")

		return api_api.API(&bootstrapApi), err

	} else {
		/**
		 * The automated build is complex enough that it deserves
		 * it's own method
		 */

		// build at least the config operations, which we will need for a config wrapper
		bootstrapApi.ActivateBuilder("local", *api_builder.New_Implementations([]string{"config", "setting", "security"}), nil)

		// Now use the config operations to determine what builds are needed
		bootstrapOps := bootstrapApi.Operations()
		bootstrapConfigWrapper := api_config.New_SimpleConfigWrapper(&bootstrapOps)

		// Use the builderConfigWrapper to list build components
		builderConfigWrapper := handlers_configwrapper.New_BuilderComponentsConfigWrapperYaml(bootstrapConfigWrapper)
		builderList := builderConfigWrapper.List()

		// this is our actual local API
		log.Debug("CLI:LocalAPI: Building LocalAPI [secured]")
		localApi := api_builder.SecureBuilderAPI{}

		built := map[string]bool{}

		for _, key := range builderList {
			builderSetting, _ := builderConfigWrapper.Get(key)
			var buildErr error

			if _, checked := built[builderSetting.Type]; !checked {
				built[builderSetting.Type] = true
				switch builderSetting.Type {
				case "local":
					log.Debug("LocalAPI: Building Local builder")
					localApi.AddBuilder(handlers_local.New_LocalBuilder(settings))
				case "upcloud":
					log.Debug("LocalAPI: Building UpCloud builder")
					localApi.AddBuilder(api_builder.Builder(&handlers_upcloud.UpcloudBuilder{}))
				default:
					buildErr = errors.New("Unrecognized builder " + builderSetting.Type)
					log.WithError(buildErr).Error("Could not build " + builderSetting.Type)
					built[builderSetting.Type] = false
				}
			}

			if success, checked := built[builderSetting.Type]; success && checked {
				log.WithFields(log.Fields{"type": builderSetting.Type, "implementations": builderSetting.Implementations.Order(), "key": key}).Debug("Activate builder from settings")
				localApi.ActivateBuilder(builderSetting.Type, builderSetting.Implementations, builderSetting.SettingsProvider)
			} else {
				log.WithError(err).WithFields(log.Fields{"builder": builderSetting.Type}).Error("Unknown builder referenced in local project")
			}
		}

		return api_api.API(&localApi), err
	}

}
