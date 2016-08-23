package command

import (
	"github.com/james-nesbitt/wundertools-go/config"

	log "github.com/Sirupsen/logrus"
)

const COMMAND_TYPE_UNKNOWN_TYPE = "unknown"

type UnknownTypeCommandSettings struct {
	Type string `yaml:"type"`
}

type UnknownTypeCommand struct {
	CommandBase
	settings   UnknownTypeCommandSettings
	persistant bool
}

func (command *UnknownTypeCommand) Init(application *config.Application) {
	command.CommandBase.Init(application)
}
func (command *UnknownTypeCommand) Settings(settings interface{}) {
	command.settings = settings.(UnknownTypeCommandSettings)
}
func (command *UnknownTypeCommand) Exec(flags ...string) {

	log.WithFields(log.Fields{"type": command.settings.Type}).Error("Unknown command type executed")

}
