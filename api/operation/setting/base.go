package setting

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * @QUESTION can't we just use a global definition for these, as they
 * are used for config and for settings (and properly others)
 */

// A setting operation base that provides a key string
type BaseSettingKeyScopeOperation struct {
	properties *operation.Properties
}

// Provides properties
func (base *BaseSettingKeyScopeOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&SettingKeyProperty{}))
		base.properties.Add(operation.Property(&SettingScopeProperty{}))
	}
	return base.properties
}

// A setting operation base that provides a string key/value property pair
type BaseSettingKeyScopeValueOperation struct {
	properties *operation.Properties
}

// Provide properties
func (base *BaseSettingKeyScopeValueOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&SettingKeyProperty{}))
		base.properties.Add(operation.Property(&SettingScopeProperty{}))
		base.properties.Add(operation.Property(&SettingValueProperty{}))
	}
	return base.properties
}

// A setting operation base that provides a key list string array
type BaseSettingKeyScopeKeysOperation struct {
	properties *operation.Properties
}

// Provide properties
func (base *BaseSettingKeyScopeKeysOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&SettingKeyProperty{}))
		base.properties.Add(operation.Property(&SettingScopeProperty{}))
		base.properties.Add(operation.Property(&SettingKeysProperty{}))
	}
	return base.properties
}
