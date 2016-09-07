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
	"os"

	log "github.com/Sirupsen/logrus"
)

// A set of Configurations
type Configurations struct {
	configurationMap   map[string]Configuration
	configurationOrder []string
}

func (configurations *Configurations) Add(configuration Configuration) {
	if configurations.configurationMap == nil {
		configurations.configurationMap = map[string]Configuration{}
		configurations.configurationOrder = []string{}
	}
	/**
	 * @TODO check if it already exists (by key)
	 */
	configurations.configurationMap[configuration.Id()] = configuration
	configurations.configurationOrder = append(configurations.configurationOrder, configuration.Id())
}
func (configurations *Configurations) Get(id string) (Configuration, bool) {
	configuration, ok := configurations.configurationMap[id]
	return configuration, ok
}
func (configurations *Configurations) Order() []string {
	return configurations.configurationOrder
}

// A single Configuration
type Configuration interface {
	// ID returns string unique configuration Identifier
	Id() string
	// Label returns a short user readable label for the configuration
	Label() string
	// Description provides a longer multi-line string description of what the Configuration does
	Description() string

	// Value allows the retrieval and setting of unknown Typed values for the Configuration.
	Get() interface{}
	Set(interface{}) bool
}

// BaseConfiguration is a Base Configuration implementation that keeps string variables for primary methods
type BaseConfiguration struct {
	id          string
	label       string
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

func (config *StringConfiguration) Get() interface{} {
	return interface{}(config.value)
}
func (config *StringConfiguration) Set(value interface{}) bool {
	if converted, ok := value.(string); ok {
		config.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Configuration value, because the passed parameter was the wrong type. Expecte string")
		return false
	}
}

// A base Configuration that provides a Bytes Array value
type BytesArrayConfiguration struct {
	value []byte
}

func (config *BytesArrayConfiguration) Get() interface{} {
	return interface{}(config.value)
}
func (config *BytesArrayConfiguration) Set(value interface{}) bool {
	if converted, ok := value.([]byte); ok {
		config.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign COnfiguration value, because the passed parameter was the wrong type. Expecte []byte")
		return false
	}
}

// A base Configuration that provides a Boolean value
type BooleanConfiguration struct {
	value bool
}

func (config *BooleanConfiguration) Get() interface{} {
	return interface{}(config.value)
}
func (config *BooleanConfiguration) Set(value interface{}) bool {
	if converted, ok := value.(bool); ok {
		config.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Configuration value, because the passed parameter was the wrong type. Expecte bool")
		return false
	}
}

// A base Configuration that provides an IO.Writer
type WriterConfiguration struct {
	value io.Writer
}

func (config *WriterConfiguration) Get() interface{} {
	if config.value == nil {
		// writer := log.StandardLogger().Writer()
		// defer writer.Close()
		// config.value = io.Writer(writer)
		config.value = io.Writer(os.Stdout)
	}

	return interface{}(config.value)
}
func (config *WriterConfiguration) Set(value interface{}) bool {
	if converted, ok := value.(io.Writer); ok {
		config.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Configuration value, because the passed parameter was the wrong type. Expecte io.Writer")
		return false
	}
}
