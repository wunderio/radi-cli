package config

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Here are the commond shared configurations for the various
 * Config operations.
 */

const (
	// config for a single config key
	OPERATION_CONFIGURATION_CONFIG_KEY = "config.key"
	// config for an orerered list of keys
	OPERATION_CONFIGURATION_CONFIG_KEYS = "config.keys"
	// config for a single config value (string)
	OPERATION_CONFIGURATION_CONFIG_VALUE = "config.value"
)

// Configuration for a single config ket
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

// Configuration for an ordered list of config keys
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

// Configuration for a single config value
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
