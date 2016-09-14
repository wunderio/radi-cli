package bytesource

import (
	"io"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
)

/**
 * File base byte stream handling
 */

// PathRoot constructor from a string path
func NewPathRoot_FromStringPath(rootPath string) *PathRoot {
	return &PathRoot{path: rootPath}
}

// A struct to hold a path route, used to find relative files in
type PathRoot struct {
	path string
}

// Get the path string for the PathRoot
func (pathRoot *PathRoot) PathString() string {
	return pathRoot.path
}

// Find a relative file path inside a PathRoot
func (pathRoot *PathRoot) FullPath(filePath string) *FileByteSource {
	return NewFileByteSource_FromPath(path.Join(pathRoot.path, filePath))
}

// Construct a FileByteSource from a file path
func NewFileByteSource_FromPath(filePath string) *FileByteSource {
	return &FileByteSource{path: filePath}
}

// A file BytesSource
type FileByteSource struct {
	path string
}

// Validate that the file exists and is readable
func (fileSource *FileByteSource) Validate() bool {
	return true
}

// Get a reader for the File
func (fileSource *FileByteSource) Reader() (io.Reader, error) {
	if osFile, err := os.Open(fileSource.path); err == nil {
		return io.Reader(osFile), err
	} else {
		log.WithFields(log.Fields{"file": fileSource.path}).WithError(err).Error("Could not make a reader from the file")
		return io.Reader(osFile), err
	}
}

// Get a reader for the File
func (fileSource *FileByteSource) Writer() (io.Writer, error) {
	// osFile, err := os.OpenFile(fileSource.path, os.O_WRONLY, os.FileMode(0666))
	osFile, err := os.Create(fileSource.path)
	return io.Writer(osFile), err
}

//
type BaseFileSourceOperation struct {
	BaseByteArraySourceOperation
	source FileByteSource
}

//
func (operation *BaseFileSourceOperation) SetFile(source FileByteSource) {
	operation.source = source
}
