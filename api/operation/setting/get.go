package setting

const (
	OPERATION_ID_SETTING_GET = "setting.get"
)

/**
 * Retrieve keyed configurations for the config handler
 */

// Base class for config get Operation
type BaseSettingGetOperation struct {
	BaseSettingKeyScopeValueOperation
}

// Id the operation
func (get *BaseSettingGetOperation) Id() string {
	return OPERATION_ID_SETTING_GET
}

// Label the operation
func (get *BaseSettingGetOperation) Label() string {
	return "Setting Get"
}

// Description for the operation
func (get *BaseSettingGetOperation) Description() string {
	return "Retrieve a keyed setting."
}

// Is this an internal API operation
func (get *BaseSettingGetOperation) Internal() bool {
	return false
}
