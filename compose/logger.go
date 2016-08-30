package compose

import (
	"errors"
	"io"

	log "github.com/Sirupsen/logrus"

	libCompose_logger "github.com/docker/libcompose/logger"
)

func NewWundertoolsLoggerFactory() libCompose_logger.Factory {
	factory := WundertoolsComposerLoggerFactory{}

	return libCompose_logger.Factory(&factory)
}

type WundertoolsComposerLoggerFactory struct {
}

func (factory *WundertoolsComposerLoggerFactory) CreateContainerLogger(name string) libCompose_logger.Logger {
	return factory.createLogger("[" + name + "]-->")
}
func (factory *WundertoolsComposerLoggerFactory) CreateBuildLogger(name string) libCompose_logger.Logger {
	return factory.createLogger("BUILD [" + name + "]-->")
}
func (factory *WundertoolsComposerLoggerFactory) CreatePullLogger(name string) libCompose_logger.Logger {
	return factory.createLogger("PULL [" + name + "]-->")
}

func (factory *WundertoolsComposerLoggerFactory) createLogger(prefix string) libCompose_logger.Logger {
	return libCompose_logger.Logger(&WundertoolsComposerLogger{
		Prefix:    prefix,
		outWriter: &WundertoolsComposerLogger_Writer{Type: "out", Prefix: prefix},
		errWriter: &WundertoolsComposerLogger_Writer{Type: "err", Prefix: prefix},
	})
}

type WundertoolsComposerLogger struct {
	Prefix    string
	outWriter *WundertoolsComposerLogger_Writer
	errWriter *WundertoolsComposerLogger_Writer
}

func (logger *WundertoolsComposerLogger) Out(bytes []byte) {
	logger.outWriter.Write(bytes)
}
func (logger *WundertoolsComposerLogger) Err(bytes []byte) {
	logger.errWriter.Write(bytes)
}

func (logger *WundertoolsComposerLogger) OutWriter() io.Writer {
	return io.Writer(logger.outWriter)
}
func (logger *WundertoolsComposerLogger) ErrWriter() io.Writer {
	return io.Writer(logger.errWriter)
}

type WundertoolsComposerLogger_Writer struct {
	Type   string
	Prefix string
}

func (writer *WundertoolsComposerLogger_Writer) Write(bytes []byte) (int, error) {
	// if the message ends in a /n, remove it
	if bytes[len(bytes)-1] == 10 {
		bytes = bytes[0 : len(bytes)-1]
	}

	message := writer.Prefix + string(bytes)

	switch writer.Type {
	case "out":
		log.Info(message)
		return len(bytes), nil
	case "err":
		log.Error(message)
		return len(bytes), nil
	}
	return 0, errors.New("wrong writer type")
}
