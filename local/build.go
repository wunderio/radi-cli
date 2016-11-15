package local

import (
	log "github.com/Sirupsen/logrus"

	api_builder "github.com/james-nesbitt/kraut-api/builder"
	api_config "github.com/james-nesbitt/kraut-api/operation/config"
	handlers_configconnect "github.com/james-nesbitt/kraut-handlers/configconnect"
	handlers_upcloud "github.com/james-nesbitt/kraut-handlers/upcloud"
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

		built := BuildTracker{}
		built.Add("local")

		for _, key := range builderList {
			builderSetting, _ := builderConfigWrapper.Get(key)

			if !built.IsBuilt(builderSetting.Type) {
				localAddBuilder(localApi, builderSetting.Type)
			}

			log.WithFields(log.Fields{"type": builderSetting.Type, "implementations": builderSetting.Implementations.Order(), "key": key}).Debug("Activate builder from settings")

			localApi.ActivateBuilder(builderSetting.Type, builderSetting.Implementations, builderSetting.Settings)
		}

	}

}

// Add builders to the BuilderAPI, considering that this is a local app
func localAddBuilder(localApi *api_builder.BuilderAPI, key string) {
	switch key {
	case "local":
		// ignore me
	case "upcloud":
		log.Debug("LocalAPI: Building UpCloud builder")
		localApi.AddBuilder(api_builder.Builder(&handlers_upcloud.UpcloudBuilder{}))
	default:
		log.WithFields(log.Fields{"builder": key}).Error("Unknown builder referenced in local project")
	}
}

// Simple struct tracks which builders have already been added
type BuildTracker struct {
	built []string
}

func (tracker *BuildTracker) Add(add string) {
	if !tracker.IsBuilt(add) {
		tracker.built = append(tracker.built, add)
	}
}
func (tracker *BuildTracker) IsBuilt(add string) bool {
	for _, built := range tracker.built {
		if add == built {
			return true
		}
	}
	return false
}
