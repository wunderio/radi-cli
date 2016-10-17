package local

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/config"
	"github.com/james-nesbitt/wundertools-go/api/operation/setting"
)

/**
 * Local handlers provides operations based entirely on the
 * Local environment, primarily based on config files in
 * a project, based on the current path'
 */

// Constructor for a localHandlerBase
func New_LocalHandler_Base(settings *LocalAPISettings) *LocalHandler_Base {
	return &LocalHandler_Base{
		settings:   settings,
		operations: &operation.Operations{},
	}
}

// A handler for base local handlers
type LocalHandler_Base struct {
	settings   *LocalAPISettings
	operations *operation.Operations
}

// Return the stored operatons
func (base *LocalHandler_Base) Operations() *operation.Operations {
	return base.operations
}

// A handler for base local handlers that use a config source (like a yml file)
type LocalHandler_ConfigWrapperBase struct {
	configWrapper config.ConfigWrapper
}

// An accessor for the ConfigBase ConfigWrapper
func (base *LocalHandler_ConfigWrapperBase) ConfigWrapper() config.ConfigWrapper {
	return base.configWrapper
}

// An accessor to set the ConfigBase ConfigWrapper
func (base *LocalHandler_ConfigWrapperBase) SetConfigWrapper(configWrapper config.ConfigWrapper) {
	base.configWrapper = configWrapper
}

// A handler for local settings
type LocalHandler_SettingWrapperBase struct {
	settingWrapper setting.SettingWrapper
}

// An accessor for the SettingBase SettingWrapper
func (base *LocalHandler_SettingWrapperBase) SettingWrapper() setting.SettingWrapper {
	return base.settingWrapper
}

// An accessor to set the SettingsBase SettingWrapper
func (base *LocalHandler_SettingWrapperBase) SetSettingWrapper(settingWrapper setting.SettingWrapper) {
	base.settingWrapper = settingWrapper
}
