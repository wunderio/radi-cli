package local

import (
	"errors"

	log "github.com/Sirupsen/logrus"

	api_api "github.com/wunderkraut/radi-api/api"
	api_builder "github.com/wunderkraut/radi-api/builder"
	api_config "github.com/wunderkraut/radi-api/operation/config"
	handlers_configwrapper "github.com/wunderkraut/radi-handlers/configwrapper"
	handlers_local "github.com/wunderkraut/radi-handlers/local"
	handlers_null "github.com/wunderkraut/radi-handlers/null"
	handlers_rancher "github.com/wunderkraut/radi-handlers/rancher"
	handlers_upcloud "github.com/wunderkraut/radi-handlers/upcloud"
)

/**
 * Build a local API, by scanning for project settings based on the
 * path.  First a number of "conf" folders are determinged, and these
 * are used to build the localAPI.
 */

// Construct a LocalAPI by checking some paths for the current user.
func MakeLocalAPI(settings handlers_local.LocalAPISettings) (api_api.API, error) {
	var err error

	/**
	 * We have local settings, from which we can build our API
	 *
	 * We will build it using a BuilderAPI, and adding the local
	 * Builder, which will be used at a minimum for a ConfigWrapper,
	 * to determine how to build the rest of the project.
	 * The reason for this preliminary step is a chicken and egg dilemna
	 * where we need an api to get project settings, but we need
	 * project settings to decide how to build an API.
	 *
	 * The resulting API will be used to bootstrap the app.
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
				case "rancher":
					log.Debug("LocalAPI: Building Rancher builder")
					localApi.AddBuilder(api_builder.Builder(&handlers_rancher.RancherBuilder{}))
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
