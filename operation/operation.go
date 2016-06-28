package operation

import (
	"github.com/james-nesbitt/wundertools-go/config"
	"github.com/james-nesbitt/wundertools-go/log"
)

func GetOperation(name string) (Operation, bool) {
	switch name {
	case "info":
		operation := Info{}
		return Operation(&operation), true
	case "compose":
		operation := Compose{}
		return Operation(&operation), true
	}
	return nil, false
}

type Operation interface {
	Init(logger log.Log, application *config.Application)
	Execute(flags ...string)
}

// Base Command class, which will receive and keep the logger, and project conf
type BaseOperation struct {
	logger      log.Log
	application *config.Application
}

// store a logger, and conf
func (operation *BaseOperation) Init(logger log.Log, application *config.Application) {
	operation.logger = logger
	operation.application = application
}
