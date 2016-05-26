package command

import (
	"github.com/james-nesbitt/wundertools-go/log"
)

type Info struct {
	BaseCommand
}

func (command *Info) Execute(flags []string) {
	logger := command.logger
	app := command.application

	logger.Message("--SETTINGS--")
	logger.Debug(log.VERBOSITY_MESSAGE, "Name:", app.Name)
	logger.Debug(log.VERBOSITY_MESSAGE, "Author:", app.Author)
	logger.Debug(log.VERBOSITY_MESSAGE, "Environment:", app.Environment)

	logger.Message("--PATHS--")
	logger.Debug(log.VERBOSITY_MESSAGE, "Conf Path keys:", app.Paths.OrderedConfPathKeys())
	logger.Debug(log.VERBOSITY_MESSAGE, "Project Paths:", app.Paths)

}
