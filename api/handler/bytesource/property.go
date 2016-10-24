package bytesource

import (
	log "github.com/Sirupsen/logrus"
)

const (
	// config for a file settings
	OPERATION_PROPERTY_BYTESOURCE_FILESETTINGS = "bytesource.filesettings"
)

/**
 * Properties which the bytesource handler uses
 */

// Project Name Property for a docker.libCompose project
type BytesourceFilesettingsProperty struct {
	value BytesourceFileSettings
}

// Id for the Property
func (filesettings *BytesourceFilesettingsProperty) Id() string {
	return OPERATION_PROPERTY_BYTESOURCE_FILESETTINGS
}

// Label for the Property
func (filesettings *BytesourceFilesettingsProperty) Label() string {
	return "Bytesource file settings"
}

// Description for the Property
func (filesettings *BytesourceFilesettingsProperty) Description() string {
	return "Filebased bytesource paths configuration object."
}

func (filesettings *BytesourceFilesettingsProperty) Get() interface{} {
	return interface{}(filesettings.value)
}
func (filesettings *BytesourceFilesettingsProperty) Set(value interface{}) bool {
	if converted, ok := value.(BytesourceFileSettings); ok {
		filesettings.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected bytesource.BytesourceFileSettings")
		return false
	}
}
