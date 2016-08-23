package command

import (
	"golang.org/x/net/context"

	libCompose_config "github.com/docker/libcompose/config"
	libCompose_project_options "github.com/docker/libcompose/project/options"

	"github.com/james-nesbitt/wundertools-go/config"
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

func (command *ContainerizedCommand) Init(application *config.Application) {
	command.CommandBase.Init(application)
}
func (command *ContainerizedCommand) Settings(settings interface{}) {
	command.settings = settings.(ContainerizedCommandSettings)
}
func (command *ContainerizedCommand) Exec(flags ...string) {

	runOptions := libCompose_project_options.Run{
		Detached: false,
	}

	// allow our app to alter the service, to do some string replacements etc
	command.application.AlterService(&command.settings.serviceConfig)

	command.project.AddConfig(command.name, &command.settings.serviceConfig)
	command.project.Run(context.Background(), command.name, flags, runOptions)

	if !command.persistant {
		deleteOptions := libCompose_project_options.Delete{
			RemoveVolume: true,
		}
		command.project.Delete(context.Background(), deleteOptions, command.name)
	}
}
