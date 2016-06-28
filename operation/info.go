package operation

import (
	"github.com/james-nesbitt/wundertools-go/log"
	"github.com/james-nesbitt/wundertools-go/compose"
)

type Info struct {
	BaseOperation
}

func (operation *Info) Execute(flags ...string) {
	logger := operation.logger
	app := operation.application

	logger.Message("--SETTINGS--")
	logger.Debug(log.VERBOSITY_MESSAGE, "Name:", app.Name)
	logger.Debug(log.VERBOSITY_MESSAGE, "Author:", app.Author)
	logger.Debug(log.VERBOSITY_MESSAGE, "Environment:", app.Environment)

	// logger.Message("--PATHS--")
	// logger.Debug(log.VERBOSITY_MESSAGE, "Conf Path keys:", app.Paths.OrderedConfPathKeys())
	// logger.Debug(log.VERBOSITY_MESSAGE, "Project Paths:", app.Paths)

	composeProject, ok := compose.MakeComposeProject(logger, app)
	if !ok {
		operation.logger.Error("could not build compose project")
		return
	}	

	composeProject.Info()

}
