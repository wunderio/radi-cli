package operation

/**
 * Command is an operation that hands off to any dynamic
 * operation handler, which are typically abstract commands
 * that are controlled using some yml in the project.
 */

import (
	"github.com/james-nesbitt/wundertools-go/command"
	"github.com/james-nesbitt/wundertools-go/config"
)

type Command struct {
	BaseOperation
	command command.Command
}

// track and Init the command
func (operation *Command) AddCommand(command command.Command) {
	operation.command = command
}

// store a logger, and conf
func (operation *Command) Init(application *config.Application) {
	operation.BaseOperation.Init(application)
	operation.command.Init(application)
}

func (operation *Command) Execute(flags ...string) {
	operation.command.Exec(flags...)
}
