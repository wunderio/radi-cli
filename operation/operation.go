package operation

import (
	"github.com/james-nesbitt/wundertools-go/command"
	"github.com/james-nesbitt/wundertools-go/config"
)

func GetOperation(application *config.Application, name string) (Operation, bool) {
	switch name {
	case "info":
		operation := Info{}
		return Operation(&operation), true
	case "compose":
		operation := Compose{}
		return Operation(&operation), true

	case "init":
		operation := Init{}
		return Operation(&operation), true
	case "init-generate":
		operation := InitGenerate{}
		return Operation(&operation), true
	}

	// dynamic operations are handled by the command functionality
	if command, ok := command.IsThisACommand(application, name); ok {
		operation := Command{}
		operation.AddCommand(command)
		return Operation(&operation), true
	}

	return nil, false
}

type Operation interface {
	Init(application *config.Application)
	Execute(flags ...string)
}

// Base Command class, which will receive and keep the project conf
type BaseOperation struct {
	application *config.Application
}

// store a conf
func (operation *BaseOperation) Init(application *config.Application) {
	operation.application = application
}
