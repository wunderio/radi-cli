package command

import (
	"github.com/james-nesbitt/wundertools-go/compose"
	"github.com/james-nesbitt/wundertools-go/config"

	libCompose_logger "github.com/docker/libcompose/logger"
)

// Determine if the string corresponds to a command name,
// and if so return a command object for it.
func IsThisACommand(application *config.Application, name string) (Command, bool) {
	commands := discoverCommands(application)
	return commands.Get(name)
}

func discoverCommands(application *config.Application) Commands {
	// initial empty map
	commands := Commands{}

	// see if there are any commands from yaml
	commands.Commands_FromYaml(application)

	return commands
}

// a map of CommandDefinitions
type Commands struct {
	commands map[string]Command
}

// list command names
func (commands *Commands) List() []string {
	list := []string{}
	for name, _ := range commands.commands {
		list = append(list, name)
	}
	return list
}

// add a new command to the list
func (commands *Commands) Add(name string, command Command) {
	if len(commands.commands) == 0 {
		commands.commands = map[string]Command{}
	}
	commands.commands[name] = command
}
func (commands *Commands) Get(name string) (Command, bool) {
	command, found := commands.commands[name]
	return command, found
}

type Command interface {
	Prepare(name string, allCommands *Commands)
	Init(app *config.Application)
	Settings(settings interface{})
	Exec(flags ...string)
}

type CommandBase struct {
	name        string
	allCommands *Commands

	application *config.Application
	project     *compose.ComposeProject
}

func (command *CommandBase) Prepare(name string, allCommands *Commands) {
	command.name = name
	command.allCommands = allCommands
}
func (command *CommandBase) Init(application *config.Application) {
	command.application = application
	command.project, _ = compose.MakeComposeProject(application, libCompose_logger.Factory(&libCompose_logger.RawLogger{}))
}
