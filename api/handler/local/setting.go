package local

import (
	"github.com/james-nesbitt/wundertools-go/api/handler/configconnect"
	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/setting"
)

// A handler for local settings
type LocalHandler_Setting struct {
	LocalHandler_Base
	LocalHandler_ConfigWrapperBase
}

// Identify the handler
func (handler *LocalHandler_Setting) Id() string {
	return "local.setting"
}

// [Handler.]Init tells the LocalHandler_Orchestrate to prepare it's operations
func (handler *LocalHandler_Setting) Init() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	ops := operation.Operations{}

	// Make a wrapper for the Settings Config interpretation, based on itnerpreting YML settings
	wrapper := configconnect.SettingsConfigWrapper(configconnect.New_BaseSettingConfigWrapperYmlOperation(handler.ConfigWrapper()))

	// Now we can add config operations that use that Base class
	ops.Add(operation.Operation(&configconnect.SettingConfigWrapperGetOperation{Wrapper: wrapper}))
	ops.Add(operation.Operation(&configconnect.SettingConfigWrapperSetOperation{Wrapper: wrapper}))
	ops.Add(operation.Operation(&configconnect.SettingConfigWrapperListOperation{Wrapper: wrapper}))

	handler.operations = &ops

	return operation.Result(&result)
}

// Make ConfigWrapper
func (handler *LocalHandler_Setting) SettingWrapper() setting.SettingWrapper {
	return setting.New_SimpleSettingWrapper(handler.operations)
}
