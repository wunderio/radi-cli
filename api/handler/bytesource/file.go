package bytesource

import (
	"io"
	"os"
	"path"
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
func (fileSource *FileByteSource) Validate() error {
	_, err := os.Stat(fileSource.path)
	return err
}

// Get a reader for the File
func (fileSource *FileByteSource) Reader() (io.Reader, error) {
	osFile, err := os.Open(fileSource.path)
	return io.Reader(osFile), err
}

// Get a reader for the File
func (fileSource *FileByteSource) Writer() (io.Writer, error) {
	// osFile, err := os.OpenFile(fileSource.path, os.O_WRONLY, os.FileMode(0666))
	osFile, err := os.Create(fileSource.path)
	return io.Writer(osFile), err
}

// An ordered set of filebytesources
type Files struct {
	fileMap map[string]*FileByteSource
	order   []string
}

// internal safe initializer
func (files *Files) safe() {
	if files.fileMap == nil {
		files.fileMap = map[string]*FileByteSource{}
		files.order = []string{}
	}
}

// Get a FileSource from the set
func (files *Files) Get(key string) (*FileByteSource, bool) {
	files.safe()

	file, found := files.fileMap[key]
	return file, found
}

// Add a FileSource to the set
func (files *Files) Add(key string, source *FileByteSource) {
	files.safe()

	if _, found := files.fileMap[key]; !found {
		files.order = append(files.order, key)
	}
	files.fileMap[key] = source
}

// Get the key order for the set
func (files *Files) Order() []string {
	files.safe()
	return files.order
}
