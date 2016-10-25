package setting

const (
	OPERATION_ID_SETTING_LIST = "setting.list"
)

/**
 * Retrieve keyed Properties for the config handler
 */

// Base class for config list Operation
type BaseSettingListOperation struct {
	BaseSettingKeyScopeKeysOperation
}

// Id the operation
func (list *BaseSettingListOperation) Id() string {
	return OPERATION_ID_SETTING_LIST
}

// Label the operation
func (list *BaseSettingListOperation) Label() string {
	return "Config List"
}

// Description for the operation
func (list *BaseSettingListOperation) Description() string {
	return "Retrieve a list of available configuration keys."
}

// Is this an internal API operation
func (list *BaseSettingListOperation) Internal() bool {
	return false
}
