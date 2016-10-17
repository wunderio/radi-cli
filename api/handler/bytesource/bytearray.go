package bytesource

import (
	"io"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

// A base operation which relies on byte
type BaseByteArraySourceOperation struct {
	source []byte
}

//
func (operation *BaseByteArraySourceOperation) FromBytes(source []byte) {
	operation.source = source
}

//
func (operation *BaseByteArraySourceOperation) FromReader(reader io.Reader) {
	source, err := ioutil.ReadAll(reader)
	if err == nil {
		operation.FromBytes(source)
	} else {
		log.WithError(err).Error("Could not read bytes from reader")
	}
}

//
func (operation *BaseByteArraySourceOperation) ToWriter(writer io.Writer) {
	writer.Write(operation.source)
}

//
func (operation *BaseByteArraySourceOperation) FromFile(fileSource *FileByteSource) {
	reader, err := fileSource.Reader()
	if err == nil {
		operation.FromReader(reader)
	} else {
		log.WithError(err).Error("Could not read bytes from file")
	}
}

//
func (operation *BaseByteArraySourceOperation) ToFile(fileSource *FileByteSource) {
	writer, err := fileSource.Writer()
	if err == nil {
		operation.ToWriter(writer)
	} else {
		log.WithError(err).Error("Could not write bytes from file")
	}
}
