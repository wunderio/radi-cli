package operation
/**
 * Command is an operation that hands off to any dynamic
 * operation handler, which are typically abstract commands
 * that are controlled using some yml in the project.
 */

import (
	"github.com/james-nesbitt/wundertools-go/config"
	"github.com/james-nesbitt/wundertools-go/log"
 	"github.com/james-nesbitt/wundertools-go/command"
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
func (operation *Command) Init(logger log.Log, application *config.Application) {
	operation.BaseOperation.Init(logger, application)
	operation.command.Init(logger, application)
}

func (operation *Command) Execute(flags ...string) {
	operation.command.Exec(flags...)
}
