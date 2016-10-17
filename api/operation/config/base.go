package config

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Base operations for configs, primarily giving some base Property structs
 */

// Config Base peration that has just a Key property
type BaseConfigKeyOperation struct {
	properties *operation.Properties
}

// Return operation properties
func (base *BaseConfigKeyOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&ConfigKeyProperty{}))
	}
	return base.properties
}

// Base Config operation that has a string key, and bytes array value property pair
type BaseConfigKeyValueOperation struct {
	properties *operation.Properties
}

// Return operation properties
func (base *BaseConfigKeyValueOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&ConfigKeyProperty{}))
		base.properties.Add(operation.Property(&ConfigValueProperty{}))
	}
	return base.properties
}

// Base Config operation that has a string key, and io.Reader value property pair
type BaseConfigKeyReadersOperation struct {
	properties *operation.Properties
}

// Return operation properties
func (base *BaseConfigKeyReadersOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&ConfigKeyProperty{}))
		base.properties.Add(operation.Property(&ConfigValueScopedReadersProperty{}))
	}
	return base.properties
}

// Base Config operation that has a string key, and io.Writer value property pair
type BaseConfigKeyWritersOperation struct {
	properties *operation.Properties
}

// Return operation properties
func (base *BaseConfigKeyWritersOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&ConfigKeyProperty{}))
		base.properties.Add(operation.Property(&ConfigValueScopedWritersProperty{}))
	}
	return base.properties
}

// Base Config operation that has a parent key, and key slice property pair
type BaseConfigKeyKeysOperation struct {
	properties *operation.Properties
}

// Return Operation properties
func (base *BaseConfigKeyKeysOperation) Properties() *operation.Properties {
	if base.properties == nil {
		base.properties = &operation.Properties{}

		base.properties.Add(operation.Property(&ConfigKeyProperty{}))
		base.properties.Add(operation.Property(&ConfigKeysProperty{}))
	}
	return base.properties
}
