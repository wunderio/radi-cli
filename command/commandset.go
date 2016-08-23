package command

import (
	"github.com/james-nesbitt/wundertools-go/config"

	log "github.com/Sirupsen/logrus"
)

const COMMAND_TYPE_SET = "set"

type CommandSetSettings struct {
	Commands []string `yaml:"commands"`
}

type CommandSet struct {
	CommandBase
	settings CommandSetSettings
}

func (command *CommandSet) Init(application *config.Application) {
	command.CommandBase.Init(application)
}
func (command *CommandSet) Settings(settings interface{}) {
	command.settings = settings.(CommandSetSettings)
}

func (command *CommandSet) Exec(flags ...string) {

	log.WithFields(log.Fields{"settings": command.settings, "flags": flags}).Debug("Running command set")

	for _, commandKey := range command.settings.Commands {

		if eachCommand, exists := command.allCommands.Get(commandKey); exists {

			eachCommand.Init(command.application)
			eachCommand.Exec(flags...)

		} else {
			log.WithFields(log.Fields{"command": commandKey}).Error("Could not find command")
		}
	}

}
