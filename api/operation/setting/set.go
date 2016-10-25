package setting

const (
	OPERATION_ID_SETTING_SET = "setting.set"
)

/**
 * Set keyed values into the Config handler
 */

// Base class for setting set Operation
type BaseSettingSetOperation struct {
	BaseSettingKeyScopeValueOperation
}

// Id the operation
func (set *BaseSettingSetOperation) Id() string {
	return OPERATION_ID_SETTING_SET
}

// Label the operation
func (set *BaseSettingSetOperation) Label() string {
	return "Config Set"
}

// Description for the operation
func (set *BaseSettingSetOperation) Description() string {
	return "Set a keyed configuration."
}

// Is this an internal API operation
func (set *BaseSettingSetOperation) Internal() bool {
	return true
}
