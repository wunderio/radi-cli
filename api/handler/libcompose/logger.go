package libcompose

import (
	"io"

	libCompose_logger "github.com/docker/libcompose/logger"
)

/**
 * Use our own implementation of the libCompose project Logger, so that
 * we can decide how and where to output the libCompose container logs
 *
 * If we don't do this, then we will not see any of the container output
 */

// libCompose.logger.Factory constructor
func NewLibcomposeLoggerFactory(outWriter, errWriter io.Writer) libCompose_logger.Factory {
	factory := LibcomposeLoggerFactory{
		outWriter: outWriter,
		errWriter: errWriter,
	}

	return libCompose_logger.Factory(&factory)
}

// libCompose.logger.Factory constructor
type LibcomposeLoggerFactory struct {
	outWriter io.Writer
	errWriter io.Writer
}

func (factory *LibcomposeLoggerFactory) CreateContainerLogger(name string) libCompose_logger.Logger {
	return factory.createLogger("[" + name + "]-->")
}
func (factory *LibcomposeLoggerFactory) CreateBuildLogger(name string) libCompose_logger.Logger {
	return factory.createLogger("BUILD [" + name + "]-->")
}
func (factory *LibcomposeLoggerFactory) CreatePullLogger(name string) libCompose_logger.Logger {
	return factory.createLogger("PULL [" + name + "]-->")
}

func (factory *LibcomposeLoggerFactory) createLogger(prefix string) libCompose_logger.Logger {
	return libCompose_logger.Logger(&LibcomposeLogger{
		outWriter: &LibcomposeLogger_Writer{Writer: factory.outWriter, Prefix: prefix},
		errWriter: &LibcomposeLogger_Writer{Writer: factory.errWriter, Prefix: prefix},
	})
}

// libCompose.logger.Logger implementation
type LibcomposeLogger struct {
	outWriter *LibcomposeLogger_Writer
	errWriter *LibcomposeLogger_Writer
}

func (logger *LibcomposeLogger) Out(bytes []byte) {
	logger.outWriter.Write(bytes)
}
func (logger *LibcomposeLogger) Err(bytes []byte) {
	logger.errWriter.Write(bytes)
}

func (logger *LibcomposeLogger) OutWriter() io.Writer {
	return io.Writer(logger.outWriter)
}
func (logger *LibcomposeLogger) ErrWriter() io.Writer {
	return io.Writer(logger.errWriter)
}

// subclass io.Writer wrapper for libCompose.logger.Logger implementation, as errors and output works the same
type LibcomposeLogger_Writer struct {
	// Final io.Writer destination
	Writer io.Writer
	// Prefix for all logging messages
	Prefix string
}

// Wrap the output writer, to add some metadata about the container as a prefix
func (writer *LibcomposeLogger_Writer) Write(bytes []byte) (int, error) {
	// if the message ends in a /n, remove it
	// if bytes[len(bytes)-1] == 10 {
	// 	bytes = bytes[0 : len(bytes)-1]
	// }

	message := writer.Prefix + string(bytes)
	n, err := writer.Writer.Write([]byte(message))
	return n, err
}
