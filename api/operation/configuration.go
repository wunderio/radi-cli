package operation

/**
 * Configurations are abstract Operations settings which 
 * are meant to be set outside of the Operation, before it
 * is executed.  Operations of a single key, from different
 * Handlers should probably try to all offer the same 
 * Configuration key/types, even if the work differently.
 *
 * This file provides a configuration collection sturct,
 * and the intercface for a single Configuration along with
 * some standard data type configuration base structs.
 *
 * A Configuration consumer should either recognized the 
 * Operation by it's keys, and then handle it's Configuraitons
 * as "knowns", or it should iterate through the Configurations
 * and use some user-interface to allow interaction.
 *
 * A Configuration is typically used, by overwriting it's
 * Value() reference with a different value.
 */

import (
	"io"
)

// A set of Configurations
type Configurations struct {
	configurationMap map[string]Configuration
	configurationOrder []string
}
func (configurations *Configurations) AddConfiguration(configuration Configuration) {
	/**
	 * @TODO check if it already exists (by key)
	 */
	configurations.configurationMap[configuration.Id()] = configuration
	configurations.configurationOrder = append( configurations.configurationOrder, configuration.Id() )
}
func (configurations *Configurations) Configuration(id string) (Configuration, bool) {
	configuration, ok := configurations.configurationMap[id]
	return configuration, ok
}
func (configurations *Configurations) ConfigurationOrder() []string {
	return configurations.configurationOrder
}

// A single Configuration
type Configuration interface {
	// ID returns string unique configuration Identifier
	Id()  string
	// Label returns a short user readable label for the configuration
	Label() string
	// Description provides a longer multi-line string description of what the Configuration does
	Description() string

	// Value allows the retrieval and setting of unknown Typed values for the Configuration.
	Value() interface{}
}

// BaseConfiguration is a Base Configuration implementation that keeps string variables for primary methods
type BaseConfiguration struct {
	id string
	label string
	description string
}
// Id returns the string id variable
func (config *BaseConfiguration) Id() string {
	return config.id
}

func (config *BaseConfiguration) Label() string {
	return config.label
}
func (config *BaseConfiguration) Description() string {
	return config.description
}

// A base Configuration that provides a String value
type StringConfiguration struct {
	value string
}
func (config *StringConfiguration) Value() interface{} {
	return interface{}(&config.value)
}

// A base Configuration that provides a Bytes Array value
type BytesArrayConfiguration struct {
	value []byte
}
func (config *BytesArrayConfiguration) Value() interface{} {
	return interface{}(&config.value)
}

// A base Configuration that provides a Boolean value
type BooleanConfiguration struct {
	value bool
}
func (config *BooleanConfiguration) Value() interface{} {
	return interface{}(&config.value)
}

// A base Configuration that provides an IO.Writer
type WriterConfiguration struct {
	value io.Writer
}
func (config *WriterConfiguration) Value() interface{} {
	return interface{}(&config.value)
}
