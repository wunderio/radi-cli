package bytesource

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Base Operation that has a BytesourceFilesettingsProperty
 * property
 */

func New_BaseBytesourceFilesettingsOperation(settings BytesourceFileSettings) *BaseBytesourceFilesettingsOperation {
	return &BaseBytesourceFilesettingsOperation{
		settings: settings,
	}
}

type BaseBytesourceFilesettingsOperation struct {
	settings   BytesourceFileSettings
	properties *operation.Properties
}

func (base *BaseBytesourceFilesettingsOperation) Properties() *operation.Properties {
	if base.properties == nil {
		settingsProp := BytesourceFilesettingsProperty{}
		settingsProp.Set(base.settings)

		base.properties = &operation.Properties{}
		base.properties.Add(operation.Property(&settingsProp))
	}
	return base.properties
}
