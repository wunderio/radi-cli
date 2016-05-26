package command

import (
	"github.com/james-nesbitt/wundertools-go/config"
	"github.com/james-nesbitt/wundertools-go/log"
)

func GetCommand(name string) (Command, bool) {
	switch name {
	case "info":
		com := Info{}
		return Command(&com), true
	case "compose":
		com := Compose{}
		return Command(&com), true
	}
	return nil, false
}

type Command interface {
	Init(logger log.Log, application *config.Application)
	Execute(flags []string)
}

// Base Command class, which will receive and keep the logger, and project conf
type BaseCommand struct {
	logger      log.Log
	application *config.Application
}

// store a logger, and conf
func (command *BaseCommand) Init(logger log.Log, application *config.Application) {
	command.logger = logger
	command.application = application
}
