package command

import (
	"golang.org/x/net/context"

	libCompose_config "github.com/docker/libcompose/config"
	libCompose_project_options "github.com/docker/libcompose/project/options"

	"github.com/james-nesbitt/wundertools-go/config"
	"github.com/james-nesbitt/wundertools-go/log"
)

const COMMAND_TYPE_CONTAINERIZED = "container"

type ContainerizedCommandSettings struct {
	serviceConfig libCompose_config.ServiceConfig
}

func (settings *ContainerizedCommandSettings) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := unmarshal(&settings.serviceConfig)
	return err
}

type ContainerizedCommand struct {
	CommandBase
	settings   ContainerizedCommandSettings
	persistant bool
}

func (command *ContainerizedCommand) Init(logger log.Log, application *config.Application) {
	command.CommandBase.Init(logger, application)
}
func (command *ContainerizedCommand) Settings(settings interface{}) {
	command.settings = settings.(ContainerizedCommandSettings)
}
func (command *ContainerizedCommand) Exec(flags ...string) {

	runOptions := libCompose_project_options.Run{
		Detached: false,
	}

	command.project.AddConfig(command.name, &command.settings.serviceConfig)
	command.project.Run(context.Background(), command.name, flags, runOptions)

	if !command.persistant {
		deleteOptions := libCompose_project_options.Delete{
			RemoveVolume: true,
		}
		command.project.Delete(context.Background(), deleteOptions, command.name)
	}
}
