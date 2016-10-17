package config

import (
	"io"
)

/**
 * Some utilities structs used to keep scoped reader/writer
 * access.  This means a list of objects which are ordered,
 * but also can be accessed by key
 */

// An ordered and keyed list of io.Readers
type ScopedReaders struct {
	scopeMap map[string]io.Reader
	order    []string
}

// internal safe initializer
func (readers *ScopedReaders) safe() {
	if readers.scopeMap == nil {
		readers.scopeMap = map[string]io.Reader{}
		readers.order = []string{}
	}
}

// Get a FileSource from the set
func (readers *ScopedReaders) Get(key string) (io.Reader, bool) {
	readers.safe()

	reader, found := readers.scopeMap[key]
	return reader, found
}

// Add a FileSource to the set
func (readers *ScopedReaders) Add(key string, source io.Reader) {
	readers.safe()

	if _, found := readers.scopeMap[key]; !found {
		readers.order = append(readers.order, key)
	}
	readers.scopeMap[key] = source
}

// Get the key order for the set
func (readers *ScopedReaders) Order() []string {
	readers.safe()
	return readers.order
}

// Ordered and keyed set of io.Writer
type ScopedWriters struct {
	scopeMap map[string]io.Writer
	order    []string
}

// internal safe initializer
func (writers *ScopedWriters) safe() {
	if writers.scopeMap == nil {
		writers.scopeMap = map[string]io.Writer{}
		writers.order = []string{}
	}
}

// Get a FileSource from the set
func (writers *ScopedWriters) Get(key string) (io.Writer, bool) {
	writers.safe()

	writer, found := writers.scopeMap[key]
	return writer, found
}

// Add a FileSource to the set
func (writers *ScopedWriters) Add(key string, source io.Writer) {
	writers.safe()

	if _, found := writers.scopeMap[key]; !found {
		writers.order = append(writers.order, key)
	}
	writers.scopeMap[key] = source
}

// Get the key order for the set
func (writers *ScopedWriters) Order() []string {
	writers.safe()
	return writers.order
}
