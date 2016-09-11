package config

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

const (
	OPERATION_CONFIGURATION_CONFIG_KEY   = "config.key"
	OPERATION_CONFIGURATION_CONFIG_KEYS  = "config.keys"
	OPERATION_CONFIGURATION_CONFIG_VALUE = "config.value"
)

// Configuration
type ConfigKeyConfiguration struct {
	operation.StringConfiguration
}

// Id for the configuration
func (confKey *ConfigKeyConfiguration) Id() string {
	return OPERATION_CONFIGURATION_CONFIG_KEY
}

// Label for the configuration
func (confKey *ConfigKeyConfiguration) Label() string {
	return "Configuration key."
}

// Description for the configuration
func (confKey *ConfigKeyConfiguration) Description() string {
	return "Configuration key."
}

// Configuration
type ConfigKeysConfiguration struct {
	operation.StringSliceConfiguration
}

// Id for the configuration
func (keyValue *ConfigKeysConfiguration) Id() string {
	return OPERATION_CONFIGURATION_CONFIG_KEYS
}

// Label for the configuration
func (keyValue *ConfigKeysConfiguration) Label() string {
	return "Configuration key list."
}

// Description for the configuration
func (keyValue *ConfigKeysConfiguration) Description() string {
	return "Configuration key list."
}

// Configuration
type ConfigValueConfiguration struct {
	operation.StringConfiguration
}

// Id for the configuration
func (confValue *ConfigValueConfiguration) Id() string {
	return OPERATION_CONFIGURATION_CONFIG_VALUE
}

// Label for the configuration
func (confValue *ConfigValueConfiguration) Label() string {
	return "Configuration value."
}

// Description for the configuration
func (confValue *ConfigValueConfiguration) Description() string {
	return "Configuration value."
}

// Configuration
type ConfigValueROConfiguration struct {
	ConfigValueConfiguration
}

// Id for the configuration
func (confValue *ConfigValueROConfiguration) ReadOnly() bool {
	return true
}

// Label for the configuration
func (confValue *ConfigValueROConfiguration) Label() string {
	return "Configuration value."
}

// Description for the configuration
func (confValue *ConfigValueROConfiguration) Description() string {
	return "Configuration value."
}
