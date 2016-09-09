package operation

import (
	"io"
	"os"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
)

/**
 * This file contains various base structs that could be used
 * to make building specific configurations easier.
 */

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

/**
 * TYPE configuraiton bases
 *
 * These Base configuration structs implement the Value accessors for
 * configurations where the value is meant to be of a specific type,
 * althought the Configuration Interface uses interface{}.
 *
 * To use them, include them in a struct that handles the other parts
 * of the interface, such as id(), label() etc.
 */

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
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Configuration value, because the passed parameter was the wrong type. Expecte []byte")
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

// A base Configuration that provides an net context
type ContextConfiguration struct {
	value context.Context
}

// Retrieve the context, or retrieve a Background context by default
func (config *ContextConfiguration) Get() interface{} {
	if config.value == nil {
		config.value = context.Background()
	}

	return interface{}(config.value)
}
func (config *ContextConfiguration) Set(value interface{}) bool {
	if converted, ok := value.(context.Context); ok {
		config.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Configuration value, because the passed parameter was the wrong type. Expected golang.org/x/net/context/Context")
		return false
	}
}
