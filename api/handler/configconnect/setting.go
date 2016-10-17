package configconnect

import (
	"errors"

	log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/setting"
)

const (
	// The Config key for settings
	CONFIG_KEY_SETTINGS = "settings"
)

/**
 * This defines what a ConfigWrapper must provide
 * to the settings operations.  This way different wrappers
 * could be used to interpret JSON or YML or whatever.
 */
type SettingsConfigWrapper interface {
	DefaultScope() string
	Get(key string) (SettingValues, bool)
	Set(key string, values SettingValues) bool
	List(parent string) []string
}

/**
 * The following 2 structs are used to keep track of settings
 * as a string map, but where each value knows from what config
 * scope it was derived.
 */

// A collection of Settings mapped by key
type Settings struct {
	valueMap map[string]SettingValues
}

// Safe initialize this struct
func (settings *Settings) safe() {
	if settings.valueMap == nil {
		settings.valueMap = map[string]SettingValues{}
	}
}

// Add a string map, for a certain scope
func (settings *Settings) MergeScope(scope string, merge map[string]string) {
	settings.safe()

	for key, value := range merge {
		var values SettingValues

		if existingValues, exists := settings.valueMap[key]; exists {
			values = existingValues
		} else {
			values = SettingValues{}
			settings.valueMap[key] = values
		}
		values.Set(scope, []byte(value))
		settings.Set(key, values)
	}
}

// Answer if the settings has no assigned values
func (settings *Settings) Set(key string, values SettingValues) {
	settings.safe()

	settings.valueMap[key] = values
}

// Answer if the settings has no assigned values
func (settings *Settings) Get(key string) (SettingValues, bool) {
	settings.safe()

	values, found := settings.valueMap[key]
	return values, found
}

// Answer if the settings has no assigned values
func (settings *Settings) Empty() bool {
	return settings.valueMap == nil
}

// Return a list of valid keys for the settings
func (settings *Settings) Keys() []string {
	settings.safe()

	keys := []string{}
	for key, _ := range settings.valueMap {
		keys = append(keys, key)
	}
	return keys
}

// Return a list of valid scopes for the settings
func (settings *Settings) Scopes() []string {
	settings.safe()

	scopes := []string{}
	for _, values := range settings.valueMap {
		for _, scope := range values.Scopes() {
			exists := false
			for _, existingScope := range scopes {
				if existingScope == scope {
					exists = true
					break
				}
			}
			if !exists {
				scopes = append(scopes, scope)
			}
		}
	}
	return scopes
}

// S single setting value, but with different values from scope
type SettingValues struct {
	settings map[string][]byte
	order    []string
}

// Safe initialize this struct
func (values *SettingValues) safe() {
	if values.settings == nil {
		values.settings = map[string][]byte{}
		values.order = []string{}
	}
}

// Merge in a settings value
func (values *SettingValues) Merge(merge SettingValues, override bool) {
	values.safe()
	for _, scope := range merge.Scopes() {
		if _, exists := values.Get(scope); !exists || override {
			scopeValue, _ := merge.Get(scope)
			values.Set(scope, scopeValue)
		}
	}
}

// Give a slice of all of the scope keys for a SettingValues
func (values *SettingValues) Scopes() []string {
	values.safe()
	scopes := []string{}
	for _, scope := range values.order {
		scopes = append(scopes, scope)
	}
	return scopes
}

// Get a settings value
func (values *SettingValues) Set(scope string, value []byte) {
	values.safe()
	values.settings[scope] = value
	values.order = append(values.order, scope)
}

// Get a settings value
func (values *SettingValues) Get(scope string) ([]byte, bool) {
	values.safe()
	if scopeValue, found := values.settings[scope]; found {
		return scopeValue, true
	} else {
		return []byte{}, false
	}
}

/**
 * Actual Operations
 */

// A Setting Get operation that uses a ConfigWrapper to retrieve values
type SettingConfigWrapperGetOperation struct {
	setting.BaseSettingGetOperation
	setting.BaseSettingKeyScopeValueOperation
	Wrapper SettingsConfigWrapper
}

// Validate the operation
func (get SettingConfigWrapperGetOperation) Validate() bool {
	return true
}

// Execute the operation
func (get SettingConfigWrapperGetOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	props := get.Properties()
	keyProp, _ := props.Get(setting.OPERATION_PROPERTY_SETTING_KEY)
	scopeProp, _ := props.Get(setting.OPERATION_PROPERTY_SETTING_SCOPE)
	valueProp, _ := props.Get(setting.OPERATION_PROPERTY_SETTING_VALUE)

	if key, ok := keyProp.Get().(string); ok {
		if value, ok := get.Wrapper.Get(key); ok {

			/**
			 * 1. look for a scope property value in the operation, and use it
			 * 2. try to look for a default scope value, and use it
			 * 3. iterate through all of the values and return the first one
			 */

			// 1. look for a scope property value
			if scope, ok := scopeProp.Get().(string); ok && scope != "" {
				if scopeValue, found := value.Get(scope); found {
					valueProp.Set(scopeValue)
				} else {
					result.Set(false, []error{errors.New("Setting connector did not find the value in the scope that you were looking for")})
				}
			} else {
				// 2. check for a default scope
				scope = get.Wrapper.DefaultScope()

				if scopeValue, found := value.Get(scope); found {
					scopeProp.Set(scope)
					valueProp.Set(scopeValue)
				} else {
					// 3. try to take the first value
					if len(value.Scopes()) > 0 {
						for _, scope := range value.Scopes() {
							scopeValue, _ := value.Get(scope)
							scopeProp.Set(scope)
							valueProp.Set(scopeValue)
							break
						}
					} else {
						result.Set(false, []error{errors.New("Setting connector did not find any value for the key that you were looking for")})
					}
				}
			}

		} else {
			log.Error("Setting connector did not find the value you were looking for")
			result.Set(false, []error{errors.New("Setting connector did not find the value you were looking for")})
		}
	} else {
		log.Error("Could not get a string value for Key from the config connector")
		result.Set(false, []error{errors.New("Could not get a string value for Key from the config connector")})
	}

	return operation.Result(&result)
}

// A Setting Set operation that uses a ConfigWrapper to assign values
type SettingConfigWrapperSetOperation struct {
	setting.BaseSettingSetOperation
	Wrapper SettingsConfigWrapper
}

// Validate the operation
func (set SettingConfigWrapperSetOperation) Validate() bool {
	return true
}

// Execute the operation
func (set SettingConfigWrapperSetOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	props := set.Properties()
	keyProp, _ := props.Get(setting.OPERATION_PROPERTY_SETTING_KEY)
	scopeProp, _ := props.Get(setting.OPERATION_PROPERTY_SETTING_SCOPE)
	valueProp, _ := props.Get(setting.OPERATION_PROPERTY_SETTING_VALUE)

	if key, okKey := keyProp.Get().(string); okKey {
		if value, okValue := valueProp.Get().([]byte); okValue {
			scope, okScope := scopeProp.Get().(string)
			if !okScope || scope == "" {
				scope = set.Wrapper.DefaultScope()
			}

			values := SettingValues{}
			values.Set(scope, value)

			if okSet := set.Wrapper.Set(key, values); !okSet {
				result.Set(false, []error{errors.New("Failed to set setting value")})
			} else {
				log.WithFields(log.Fields{"key": okKey, "scope": scope, "values": values}).Debug("Set config value")
			}
		} else {
			result.Set(false, []error{errors.New("Could not retrieve Value property for setting Set operation. No value to set.")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not assign value to key property for setting Set operation")})
	}

	return operation.Result(&result)
}

//A setting List operation that uses a ConfigWrapper to list keys
type SettingConfigWrapperListOperation struct {
	setting.BaseSettingListOperation
	setting.BaseSettingKeyScopeKeysOperation
	Wrapper SettingsConfigWrapper
}

// Validate the operation
func (list SettingConfigWrapperListOperation) Validate() bool {
	return true
}

// Execute the operation
func (list SettingConfigWrapperListOperation) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	props := list.BaseSettingKeyScopeKeysOperation.Properties()
	keyProp, _ := props.Get(setting.OPERATION_PROPERTY_SETTING_KEY)
	keysConf, _ := props.Get(setting.OPERATION_PROPERTY_SETTING_KEYS)

	if key, ok := keyProp.Get().(string); ok && key != "" {
		keysConf.Set(list.Wrapper.List(key))
	} else {
		keysConf.Set(list.Wrapper.List(""))
	}

	return operation.Result(&result)
}
