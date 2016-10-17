package config

import (
	"errors"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * A wrapper for config operations
 *
 * @NOTE this is currently a blocking inline process, which would stall
 *   if the backend operations timeout.  A thread-safe implementation should
 *   be written, but we should see it in operation before we do that.
 *
 * @TODO Make this much more intelligent, right now it is just a quick operator
 */

// Constructor for SimpleConfigWrapper
func New_SimpleConfigWrapper(operations *operation.Operations) *SimpleConfigWrapper {
	return &SimpleConfigWrapper{operations: operations}
}

// Simple config wrapper
type SimpleConfigWrapper struct {
	operations *operation.Operations
}

// Perform the Get Operation
func (wrapper *SimpleConfigWrapper) Get(key string) (ConfigScopedValues, error) {
	var found, success bool
	var op operation.Operation
	var keyProp, readersProp operation.Property

	result := ConfigScopedValues{}

	if op, found = wrapper.operations.Get(OPERATION_ID_CONFIG_READERS); !found {
		return result, errors.New("No get operation available in Config Single Wrapper")
	}

	if keyProp, found = op.Properties().Get(OPERATION_PROPERTY_CONFIG_KEY); !found {
		return result, errors.New("No key configuraiton available in Config Single Wrapper")
	}

	if !keyProp.Set(key) {
		return result, errors.New("Key property value failed to set in Config Single Wrapper")
	}

	if readersProp, found = op.Properties().Get(OPERATION_PROPERTY_CONFIG_VALUE_READERS); !found {
		return result, errors.New("No value property available in Config Single Wrapper")
	}

	if success, _ = op.Exec().Success(); !success {
		return result, errors.New("Operation failed to execute in Config Single Wrapper")
	}

	readers := readersProp.Get().(ScopedReaders)
	for _, scope := range readers.Order() {
		reader, _ := readers.Get(scope)
		if contents, err := ioutil.ReadAll(reader); err == nil {
			result.Add(scope, ConfigScopedValue(contents))
		}
	}
	return result, nil
}

// Perform the Set Operation
func (wrapper *SimpleConfigWrapper) Set(key string, values ConfigScopedValues) error {
	var found, success bool
	var op operation.Operation
	var keyProp, writersProp operation.Property

	if op, found = wrapper.operations.Get(OPERATION_ID_CONFIG_WRITERS); !found {
		return errors.New("No get operation available in Config Single Wrapper")
	}

	if keyProp, found = op.Properties().Get(OPERATION_PROPERTY_CONFIG_KEY); !found {
		return errors.New("No key configuraiton available in Config Single Wrapper")
	}

	if !keyProp.Set(key) {
		return errors.New("Key property value failed to set in Config Single Wrapper")
	}

	if writersProp, found = op.Properties().Get(OPERATION_PROPERTY_CONFIG_VALUE_WRITERS); !found {
		return errors.New("No writers property available in Config Single Wrapper")
	}

	if success, _ = op.Exec().Success(); !success {
		return errors.New("Operation failed to execute in Config Single Wrapper")
	}

	writers := writersProp.Get().(ScopedWriters)

	var returnError error
	for _, scope := range values.Order() {
		content, _ := values.Get(scope)
		if writer, found := writers.Get(scope); found {
			log.WithFields(log.Fields{"scope": scope, "content": string(content)}).Debug("ConfigWrapper: writing config to writer")
			byteContent := []byte(content)
			if _, err := writer.Write(byteContent); err != nil {
				returnError = err
			}
		} else {
			log.WithFields(log.Fields{"scope": scope, "scopes": writers.Order()}).Error("ConfigWrapper could not find wrapper for targeted Config Set")
		}

		/**
		 * @TODO should we allow an attempt to create a new writer?
		 */
	}

	return returnError
}

// Performe the List Operation
func (wrapper *SimpleConfigWrapper) List(parent string) ([]string, error) {
	var found, success bool
	var op operation.Operation
	var keyProp, keysProp operation.Property
	var errs []error

	result := []string{}

	if op, found = wrapper.operations.Get(OPERATION_ID_CONFIG_LIST); !found {
		return result, errors.New("No list operation available in Config Wrapper")
	}

	if keyProp, found = op.Properties().Get(OPERATION_PROPERTY_CONFIG_KEY); !found {
		return result, errors.New("No key property available in Config Wrapper")
	}

	if !keyProp.Set(parent) {
		return result, errors.New("Key property value failed to set in Config Wrapper")
	}

	if keysProp, found = op.Properties().Get(OPERATION_PROPERTY_CONFIG_KEYS); !found {
		return result, errors.New("No keys property available in Config Wrapper")
	}

	if success, errs = op.Exec().Success(); !success {
		return result, errs[0]
	}

	result = keysProp.Get().([]string)
	return result, nil
}
