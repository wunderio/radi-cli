package local

import (
	"errors"

	log "github.com/Sirupsen/logrus"

	api_builder "github.com/james-nesbitt/kraut-api/builder"
	api_config "github.com/james-nesbitt/kraut-api/operation/config"
	handlers_configwrapper "github.com/james-nesbitt/kraut-handlers/configwrapper"
	handlers_upcloud "github.com/james-nesbitt/kraut-handlers/upcloud"
)

// Take a LocalAPI, and add builders to it based on what can be found from the config operations
/**
 * This function uses an APIs existing config operations, to act as a source of build information
 * which is used to add builders to the API.  The build process tries to identify builders based
 * on the config Type: information,  through simple string matching.
 */
func LocalBuild(localApi *api_builder.BuilderAPI) {

	// build at least the config operations, which we will need for a config wrapper
	localApi.ActivateBuilder("local", *api_builder.New_Implementations([]string{"config"}), nil)

	// Now use the config operations to determine what builds are needed
	configOps := localApi.Operations()
	configWrapper := api_config.New_SimpleConfigWrapper(&configOps)

	// Use the builderConfigWrapper to list build components
	builderConfigWrapper := handlers_configwrapper.New_BuilderSettingsConfigWrapperYaml(configWrapper)
	builderList := builderConfigWrapper.List()

	if len(builderList) == 0 {

		// build all local operations, as no definitions were found
		localApi.ActivateBuilder("local", *api_builder.New_Implementations([]string{"config", "setting", "project", "orchestrate", "command"}), nil)

	} else {

		built := map[string]bool{}
		built["local"] = true

		for _, key := range builderList {
			builderSetting, _ := builderConfigWrapper.Get(key)
			var err error

			if _, checked := built[builderSetting.Type]; !checked {
				if err = localAddBuilder(localApi, builderSetting.Type); err == nil {
					built[builderSetting.Type] = true
				} else {
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

	}

}

// Add builders to the BuilderAPI, so that they can be activated
/**
 * This function needs to be aware of what builders are allowed in thie
 * localAPI, which means that it is tied to the .go import() items.
 */
func localAddBuilder(localApi *api_builder.BuilderAPI, key string) error {
	switch key {
	case "local":
		// ignore me
	case "upcloud":
		log.Debug("LocalAPI: Building UpCloud builder")
		localApi.AddBuilder(api_builder.Builder(&handlers_upcloud.UpcloudBuilder{}))
	default:
		return errors.New("Unrecognized builder " + key)
	}
	return nil
}
