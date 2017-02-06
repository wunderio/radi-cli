package local

import (
	"errors"

	log "github.com/Sirupsen/logrus"

	api_api "github.com/wunderkraut/radi-api/api"
	api_builder "github.com/wunderkraut/radi-api/builder"
	api_config "github.com/wunderkraut/radi-api/operation/config"
	handler_rancher "github.com/wunderkraut/radi-handler-rancher"
	handler_upcloud "github.com/wunderkraut/radi-handler-upcloud"
	handler_configwrapper "github.com/wunderkraut/radi-handlers/configwrapper"
	handler_local "github.com/wunderkraut/radi-handlers/local"
	handler_null "github.com/wunderkraut/radi-handlers/null"
)

/**
 * Build a local API, by scanning for project settings based on the
 * path.  First a number of "conf" folders are determinged, and these
 * are used to build the localAPI.
 */

// Construct a LocalAPI by checking some paths for the current user.
func MakeLocalAPI_LocalSecure(settings handler_local.LocalAPISettings, configWrapper api_config.ConfigWrapper) (api_api.API, error) {
	var err error

	// this is our actual local API
	log.Debug("CLI:LocalAPI: Building LocalAPI [secured]")
	localApi := api_builder.SecureBuilderAPI{}

	// Use the builderConfigWrapper to list build components
	log.Debug("CLI:LocalAPI: Building LocalAPI: Building builder configwrapper")
	builderConfigWrapper := handler_configwrapper.New_BuilderComponentsConfigWrapperYaml(configWrapper)
	builderList := builderConfigWrapper.List()
	log.WithFields(log.Fields{"list": builderList}).Debug("CLI:LocalAPI: Building LocalAPI: Built builder configwrapper")

	built := map[string]bool{}

	for _, key := range builderList {
		builderSetting, _ := builderConfigWrapper.Get(key)
		var buildErr error

		if _, checked := built[builderSetting.Type]; !checked {
			log.WithFields(log.Fields{"builder": builderSetting.Type}).Debug("CLI:LocalAPI: Building LocalAPI : AddingBuilder")
			built[builderSetting.Type] = true
			switch builderSetting.Type {
			case "null":
				log.Debug("LocalAPI: Building Null builder")
				localApi.AddBuilder(handler_null.New_NullBuilder())
			case "local":
				log.Debug("LocalAPI: Building Local builder")
				localApi.AddBuilder(handler_local.New_LocalBuilder(settings))
			case "upcloud":
				log.Debug("LocalAPI: Building UpCloud builder")
				localApi.AddBuilder(api_builder.Builder(&handler_upcloud.UpcloudBuilder{}))
			case "rancher":
				log.Debug("LocalAPI: Building Rancher builder")
				localApi.AddBuilder(api_builder.Builder(&handler_rancher.RancherBuilder{}))
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
