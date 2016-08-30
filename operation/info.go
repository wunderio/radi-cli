package operation

import (
	log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/compose"

	libCompose_logger "github.com/docker/libcompose/logger"
)

type Info struct {
	BaseOperation
}

func (operation *Info) Execute(flags ...string) {
	app := operation.application

	log.WithFields(log.Fields{
		"Name":        app.Name,
		"Author":      app.Author,
		"Environment": app.Environment,
	}).Info("Settings")

	// logger.Message("--PATHS--")
	// logger.Debug(log.VERBOSITY_MESSAGE, "Conf Path keys:", app.Paths.OrderedConfPathKeys())
	// logger.Debug(log.VERBOSITY_MESSAGE, "Project Paths:", app.Paths)

	composeProject, ok := compose.MakeComposeProject(app, libCompose_logger.Factory(&libCompose_logger.RawLogger{}))
	if !ok {
		log.Error("could not build compose project")
		return
	}

	composeProject.Info()
}
