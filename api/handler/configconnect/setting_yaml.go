package configconnect

import (
	// "errors"
	"strings"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/james-nesbitt/wundertools-go/api/operation/config"
)

/**
 * A single base struct which handles settings by intepreting a config.ConfigWrapper
 * as a stream of YML bytes,
 */

// Constructor for {} BaseSettingConfigWrapperYmlOperation
func New_BaseSettingConfigWrapperYmlOperation(wrapper config.ConfigWrapper) *BaseSettingConfigWrapperYmlOperation {
	return &BaseSettingConfigWrapperYmlOperation{
		wrapper:  wrapper,
		settings: Settings{},
	}
}

// A SettingsSource implementation for yml settings
type BaseSettingConfigWrapperYmlOperation struct {
	wrapper  config.ConfigWrapper // The config wrapper will be used to retrieve and save full config
	settings Settings             // the values map stores parsed values from config
}

// Retrieve values by parsing bytes from the wrapper
func (setting *BaseSettingConfigWrapperYmlOperation) Load() error {
	setting.settings = Settings{} // reset stored settings so that we can repopulate it.
	if sources, err := setting.wrapper.Get(CONFIG_KEY_SETTINGS); err == nil {
		for _, scope := range sources.Order() {
			scopedSource, _ := sources.Get(scope)
			scopedValues := map[string]string{} // temporarily hold all settings for a specific scope in this
			if err := yaml.Unmarshal(scopedSource, &scopedValues); err == nil {
				setting.settings.MergeScope(scope, scopedValues)
			} else {
				log.WithError(err).WithFields(log.Fields{"scope": scope}).Error("Couldn't marshall yml scope")
			}
			log.WithFields(log.Fields{"bytes": string(scopedSource), "values": scopedValues, "settings": setting}).Debug("Settings:Config->Load()")
		}
		return nil
	} else {
		log.WithError(err).Error("Error loading config for " + CONFIG_KEY_SETTINGS)
		return err
	}
}

// Save the current values to the wrapper
func (setting *BaseSettingConfigWrapperYmlOperation) Save() error {
	// create and initialize some primitve map for holding all settings by scope
	configMap := map[string]map[string]string{} // map[scope]map[key]value
	for _, scope := range setting.settings.Scopes() {
		configMap[scope] = map[string]string{}
	}

	// Map all of the scoped values into the map
	for _, key := range setting.settings.Keys() {
		scopedValues, _ := setting.settings.Get(key)

		for _, scope := range scopedValues.Scopes() {
			scopedValue, _ := scopedValues.Get(scope)

			configMap[scope][key] = string(scopedValue)
		}
	}

	// convert the map to a ConfigScopedValues{} by marshalling the settings maps
	scopedValues := config.ConfigScopedValues{}
	for scope, values := range configMap {
		if valuesYml, err := yaml.Marshal(values); err == nil {
			scopedValues.Add(scope, config.ConfigScopedValue(valuesYml))
		} else {
			return err
		}
	}

	// Use the Config wrapper to save the scoped values
	setting.wrapper.Set(CONFIG_KEY_SETTINGS, scopedValues)

	return nil
}

// Return the default scope string for the wrapper
func (setting *BaseSettingConfigWrapperYmlOperation) DefaultScope() string {
	/**
	 * @TODO come up with better scopes, but it has to match local conf path keys
	 */
	return "project-wundertools"
}

// SettingSource interface List implementation
func (setting *BaseSettingConfigWrapperYmlOperation) Get(key string) (SettingValues, bool) {
	if setting.settings.Empty() {
		setting.Load()
	}
	value, found := setting.settings.Get(key)

	log.WithFields(log.Fields{"key": key, "value": value, "found": found, "settings": setting}).Debug("Settings:Config->Get()")
	return value, found
}

// SettingSource interface List implementation
func (setting *BaseSettingConfigWrapperYmlOperation) Set(key string, values SettingValues) bool {
	if setting.settings.Empty() {
		setting.Load()
	}

	setting.settings.Set(key, values)
	if err := setting.Save(); err == nil {
		return true
	} else {
		log.WithError(err).Error("Could not set setting, Config wrapper failed to save")
		return false
	}
}

// SettingSource interface List implementation
func (setting *BaseSettingConfigWrapperYmlOperation) List(parent string) []string {
	if setting.settings.Empty() {
		setting.Load()
	}

	keys := []string{}
	for _, key := range setting.settings.Keys() {
		if parent == "" || (key != parent && strings.HasPrefix(key, parent)) {
			keys = append(keys, key)
		}
	}
	return keys
}
