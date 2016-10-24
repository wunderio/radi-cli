package operation

import (
	"io"
	"os"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
)

/**
 * This file contains various base structs that could be used
 * to make building specific operation properties easier.
 */

// BaseProperty is A base Property implementation that keeps string variables for primary methods
type BaseProperty struct {
	id          string
	label       string
	description string
	internal    bool
}

// Id returns the string id variable
func (property *BaseProperty) Id() string {
	return property.id
}

func (property *BaseProperty) Label() string {
	return property.label
}
func (property *BaseProperty) Description() string {
	return property.description
}
func (property *BaseProperty) Internal() bool {
	return property.internal
}

/**
 * TYPE property bases
 *
 * These Base property structs implement the Value accessors for
 * properties where the value is meant to be of a specific type,
 * although the Property Interface uses interface{}.  This makes
 * it easier to implement a property that just tracks a string,
 * bool or maybe an io.Reader type.
 *
 * To use them, include them in a struct that handles the other parts
 * of the interface, such as id(), label() etc.
 */

// A base Property that provides a String value
type StringProperty struct {
	value string
}

// Give an idea of what type of value the property consumes
func (property *StringProperty) Type() string {
	return "string"
}

func (property *StringProperty) Get() interface{} {
	return interface{}(property.value)
}
func (property *StringProperty) Set(value interface{}) bool {
	if converted, ok := value.(string); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected string")
		return false
	}
}

// A base Property that provides a slice of string values
type StringSliceProperty struct {
	value []string
}

// Give an idea of what type of value the property consumes
func (property *StringSliceProperty) Type() string {
	return "[]string"
}

func (property *StringSliceProperty) Get() interface{} {
	return interface{}(property.value)
}
func (property *StringSliceProperty) Set(value interface{}) bool {
	if converted, ok := value.([]string); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected []string")
		return false
	}
}

// A base Property that provides a Bytes Array value
type BytesArrayProperty struct {
	value []byte
}

// Give an idea of what type of value the property consumes
func (property *BytesArrayProperty) Type() string {
	return "[]byte"
}

func (property *BytesArrayProperty) Get() interface{} {
	return interface{}(property.value)
}
func (property *BytesArrayProperty) Set(value interface{}) bool {
	if converted, ok := value.([]byte); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected []byte")
		return false
	}
}

// A base Property that provides a Boolean value
type BooleanProperty struct {
	value bool
}

// Give an idea of what type of value the property consumes
func (property *BooleanProperty) Type() string {
	return "bool"
}

func (property *BooleanProperty) Get() interface{} {
	return interface{}(property.value)
}
func (property *BooleanProperty) Set(value interface{}) bool {
	if converted, ok := value.(bool); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected bool")
		return false
	}
}

// A base Property that provides an IO.Writer
type WriterProperty struct {
	value io.Writer
}

// Give an idea of what type of value the property consumes
func (property *WriterProperty) Type() string {
	return "io.Writer"
}

func (property *WriterProperty) Get() interface{} {
	if property.value == nil {
		// writer := log.StandardLogger().Writer()
		// defer writer.Close()
		// property.value = io.Writer(writer)
		property.value = io.Writer(os.Stdout)
	}
	return interface{}(property.value)
}
func (property *WriterProperty) Set(value interface{}) bool {
	if converted, ok := value.(io.Writer); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected io.Writer")
		return false
	}
}

// A base Property that provides an IO.Reader
type ReaderProperty struct {
	value io.Reader
}

// Give an idea of what type of value the property consumes
func (property *ReaderProperty) Type() string {
	return "io.Reader"
}

func (property *ReaderProperty) Get() interface{} {
	if property.value == nil {
		property.value = io.Reader(os.Stdin)
	}
	return interface{}(property.value)
}
func (property *ReaderProperty) Set(value interface{}) bool {
	if converted, ok := value.(io.Reader); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected io.Reader")
		return false
	}
}

// A base Property that provides an net context
type ContextProperty struct {
	value context.Context
}

// Give an idea of what type of value the property consumes
func (property *ContextProperty) Type() string {
	return "golang.org/x/net/context.Context"
}

// Retrieve the context, or retrieve a Background context by default
func (property *ContextProperty) Get() interface{} {
	if property.value == nil {
		property.value = context.Background()
	}

	return interface{}(property.value)
}
func (property *ContextProperty) Set(value interface{}) bool {
	if converted, ok := value.(context.Context); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected golang.org/x/net/context/Context")
		return false
	}
}
