package libcompose

import (
	"errors"
	"strings"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"

	libCompose_config "github.com/docker/libcompose/config"
	libCompose_project_options "github.com/docker/libcompose/project/options"

	"github.com/james-nesbitt/wundertools-go/api/operation"
	"github.com/james-nesbitt/wundertools-go/api/operation/command"
	"github.com/james-nesbitt/wundertools-go/api/operation/config"
)

// Constructor for BaseCommandConfigWrapperYmlOperation
func New_BaseCommandConfigWrapperYmlOperation(configWrapper config.ConfigWrapper) *BaseCommandConfigWrapperYmlOperation {
	return &BaseCommandConfigWrapperYmlOperation{
		wrapper:  configWrapper,
		commands: &CommandYmlCommands{},
	}
}

// Command config wrapper that reads YML commands
type BaseCommandConfigWrapperYmlOperation struct {
	wrapper  config.ConfigWrapper
	commands *CommandYmlCommands
}

// Load all commands from config
func (commands *BaseCommandConfigWrapperYmlOperation) load() error {
	commands.commands = &CommandYmlCommands{} // reset the command list
	if sources, err := commands.wrapper.Get(CONFIG_KEY_COMMAND); err == nil {
		for _, scope := range sources.Order() {
			scopedSource, _ := sources.Get(scope)
			scopedCommands := CommandYmlCommands{} // temporarily hold all commands for a specific scope in this
			if err := yaml.Unmarshal(scopedSource, &scopedCommands); err == nil {
				commands.commands.Merge(scopedCommands)
				log.WithFields(log.Fields{"scope": scope, "merged": commands.commands.Order(), "new": scopedCommands.Order()}).Debug("Commands:Config->Load()")
			} else {
				log.WithError(err).WithFields(log.Fields{"scope": scope}).Error("Couldn't marshall yml scope")
			}
		}
		return nil
	} else {
		log.WithError(err).Error("Error loading config for " + CONFIG_KEY_COMMAND)
		return err
	}
} // Save Commands
func (commands *BaseCommandConfigWrapperYmlOperation) save() error {
	return errors.New("Not allowed to save commands yet")
}

// CommandSource interface List implementation
func (commands *BaseCommandConfigWrapperYmlOperation) Get(key string) (*CommandYmlCommand, error) {
	if commands.commands.Empty() {
		commands.load()
	}
	if comm, err := commands.commands.Get(key); err == nil {
		log.WithFields(log.Fields{"key": key, "comm": comm}).WithError(err).Debug("Commands:Config->Get()")
		return comm, nil
	} else {
		return nil, err
	}
}

// CommandSource interface List implementation
func (commands *BaseCommandConfigWrapperYmlOperation) Set(key string, comm *CommandYmlCommand) error {
	return errors.New("Not yet able to set Commands")
}

// CommandSource interface List implementation
func (commands *BaseCommandConfigWrapperYmlOperation) List(parent string) ([]string, error) {
	if commands.commands.Empty() {
		commands.load()
	}

	keys := []string{}
	for _, key := range commands.commands.Order() {
		if parent == "" || (key != parent && strings.HasPrefix(key, parent)) {
			keys = append(keys, key)
		}
	}
	return keys, nil
}

type CommandYmlCommands struct {
	comms map[string]*CommandYmlCommand
	order []string
}

// Yaml UnMarshaller
func (comms *CommandYmlCommands) UnmarshalYAML(unmarshal func(interface{}) error) error {
	holder := map[string]*CommandYmlCommand{}
	if error := unmarshal(&holder); error == nil {
		for key, comm := range holder {
			comm.setId(key)
			comms.Set(key, comm)
		}
		return nil
	} else {
		return error
	}
}

// Safe lazy constructor
func (comms *CommandYmlCommands) safe() {
	if comms.comms == nil {
		comms.comms = map[string]*CommandYmlCommand{}
		comms.order = []string{}
	}
}

// Safe lazy constructor
func (comms *CommandYmlCommands) Empty() bool {
	return (&comms.comms == nil) || (len(comms.comms) == 0)
}

// Add a command
func (comms *CommandYmlCommands) Set(key string, comm *CommandYmlCommand) error {
	comms.safe()
	if _, exists := comms.comms[key]; !exists {
		comms.order = append(comms.order, key)
	}
	comms.comms[key] = comm
	return nil
}

// Get a comm
func (comms *CommandYmlCommands) Get(key string) (*CommandYmlCommand, error) {
	comms.safe()
	if com, found := comms.comms[key]; found {
		return com, nil
	} else {
		return com, errors.New("Command not found")
	}
}

// Comm order
func (comms *CommandYmlCommands) Order() []string {
	comms.safe()
	return comms.order
}

// Comm merge
func (comms *CommandYmlCommands) Merge(merge CommandYmlCommands) error {
	comms.safe()
	for _, key := range merge.Order() {
		if _, err := comms.Get(key); err != nil {
			mergeComm, _ := merge.Get(key)
			comms.Set(key, mergeComm)
		}
	}
	return nil
}

// A  struct to hold yml commands, which can be used to create command.Command objects
type CommandYmlCommand struct {
	scope string
	id    string

	label       string
	description string

	persistant bool
	internal   bool

	project       *ComposeProject
	properties    *operation.Properties
	serviceConfig libCompose_config.ServiceConfig
}

// Yaml UnMarshaller
func (comm *CommandYmlCommand) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var holder struct {
		Scope       string
		Id          string
		Label       string
		Description string

		Persistant bool
		Internal   bool
	}
	if error := unmarshal(&holder); error == nil {
		comm.id = holder.Id
		comm.description = holder.Description
		comm.scope = holder.Scope
		comm.persistant = holder.Persistant
		comm.internal = holder.Internal
	}

	var serviceHolder libCompose_config.ServiceConfig
	if error := unmarshal(&serviceHolder); error == nil {
		comm.serviceConfig = serviceHolder
	}

	if comm.properties == nil {
		properties := operation.Properties{}

		properties.Add(operation.Property(&command.CommandFlagsProperty{}))

		comm.properties = &properties
	}

	return nil
}

// Turn this CommandYmlCommand into a command.Command
func (ymlCommand *CommandYmlCommand) Command(project *ComposeProject) command.Command {
	ymlCommand.project = project
	return command.Command(ymlCommand)
}

// Return string Id
func (ymlCommand *CommandYmlCommand) setId(id string) {
	ymlCommand.id = id
}

// Return string Scope
func (ymlCommand *CommandYmlCommand) Scope() string {
	return ymlCommand.scope
}

/**
 * Command interace
 */

func (ymlCommand *CommandYmlCommand) Validate() bool {
	return true
}

func (ymlCommand *CommandYmlCommand) Internal() bool {
	return ymlCommand.internal
}

// Return string Id
func (ymlCommand *CommandYmlCommand) Id() string {
	return ymlCommand.id
}

// Return string Label
func (ymlCommand *CommandYmlCommand) Label() string {
	return ymlCommand.label
}

// Return string Description
func (ymlCommand *CommandYmlCommand) Description() string {
	return ymlCommand.description
}

// Return string Description
func (ymlCommand *CommandYmlCommand) Properties() *operation.Properties {
	return ymlCommand.properties
}

func (ymlCommand *CommandYmlCommand) Exec() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	flags := []string{}
	if propFlags, found := ymlCommand.Properties().Get(command.OPERATION_PROPERTY_COMMAND_FLAGS); found {
		flags = propFlags.Get().([]string)
	}

	// @TODO GET this from a property
	runContext := context.Background()

	runOptions := libCompose_project_options.Run{
		Detached: false,
	}

	// get the service for the command
	service := ymlCommand.serviceConfig

	// allow our app to alter the service, to do some string replacements etc
	ymlCommand.project.AlterService(&service)

	ymlCommand.project.AddConfig(ymlCommand.Id(), &service)
	ymlCommand.project.Run(runContext, ymlCommand.Id(), flags, runOptions)

	if !ymlCommand.persistant {
		deleteOptions := libCompose_project_options.Delete{
			RemoveVolume: true,
		}
		ymlCommand.project.Delete(runContext, deleteOptions, ymlCommand.Id())
	}

	return operation.Result(&result)
}
