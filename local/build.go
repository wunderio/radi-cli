package local

import (
	log "github.com/Sirupsen/logrus"

	api_builder "github.com/james-nesbitt/kraut-api/builder"
	api_config "github.com/james-nesbitt/kraut-api/operation/config"
	handlers_configconnect "github.com/james-nesbitt/kraut-handlers/configconnect"
)

func LocalBuild(localApi *api_builder.BuilderAPI) {

	// build at least the config operations, which we will need for a config wrapper
	localApi.ActivateBuilder("local", *api_builder.New_Implementations([]string{"config"}), nil)

	// Now use the config operations to determine what builds are needed
	configOps := localApi.Operations()
	configWrapper := api_config.New_SimpleConfigWrapper(&configOps)

	// Use the builderConfigWrapper to list build components
	builderConfigWrapper := handlers_configconnect.New_BuilderSettingsConfigWrapperYaml(configWrapper)
	builderList := builderConfigWrapper.List()

	if len(builderList) == 0 {

		// build all local operations, as no definitions were found
		localApi.ActivateBuilder("local", *api_builder.New_Implementations([]string{"config", "setting", "project", "orchestrate", "command"}), nil)

	} else {

		for _, key := range builderList {
			builderSetting, _ := builderConfigWrapper.Get(key)

			log.WithFields(log.Fields{"type": builderSetting.Type, "implementations": builderSetting.Implementations.Order(), "key": key}).Debug("Activate builder from settings")

			localApi.ActivateBuilder(builderSetting.Type, builderSetting.Implementations, builderSetting.Settings)
		}

	}

}
