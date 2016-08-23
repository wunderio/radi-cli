package command

import (
	"io/ioutil"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/james-nesbitt/wundertools-go/config"
)

const (
	WUNDERTOOLS_COMMAND_YAML_FILE = "commands.yml"
)

// iterate through all of the conf paths, and load commands.yml
func (commands *Commands) Commands_FromYaml(application *config.Application) {

	for _, key := range application.Paths.OrderedConfPathKeys() {
		confPath, _ := application.Paths.Path(key)
		yamlFilePath := path.Join(confPath, WUNDERTOOLS_COMMAND_YAML_FILE)

		if _, err := os.Stat(yamlFilePath); err == nil {
			yamlFile, err := ioutil.ReadFile(yamlFilePath)
			if err == nil {
				commands.from_ConfYamlBytes(yamlFile)
			}
		}
	}

}

/**
 * configure an command from a yaml stream of Bytes
 * @TODO make this a reader?
 */
func (commands *Commands) from_ConfYamlBytes(yamlBytes []byte) {
	// parse the config file contents as a ConfSource_projectyaml object
	source := new(CommandsFromYaml)
	if err := yaml.Unmarshal(yamlBytes, source); err == nil {

		for name, yamlCommand := range *source {

			yamlCommand.command.Prepare(name, commands)
			commands.Add(name, yamlCommand.command)

		}

	} else {
		log.WithError(err).Warn("Could not parse commands yml.")
	}
}

type CommandsFromYaml map[string]CommandFromYaml

type CommandFromYaml struct {
	command Command
}

func (command *CommandFromYaml) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var asAMap map[string]interface{}
	unmarshal(&asAMap)

	var err error

	if typeVal, hasType := asAMap["type"]; hasType {
		if typeString, isTypeString := typeVal.(string); isTypeString {

			switch typeString {
			case COMMAND_TYPE_CONTAINERIZED:

				settings := ContainerizedCommandSettings{}
				if err = unmarshal(&settings); err == nil {

					command.command = Command(&ContainerizedCommand{})
					command.command.Settings(settings)

				}
			case COMMAND_TYPE_EXEC:

				settings := ExecCommandSettings{}
				if err = unmarshal(&settings); err == nil {

					command.command = Command(&ExecCommand{})
					command.command.Settings(settings)

				}
			case COMMAND_TYPE_SET:

				settings := CommandSetSettings{}
				if err = unmarshal(&settings); err == nil {

					command.command = Command(&CommandSet{})
					command.command.Settings(settings)

				}

			default:

				settings := UnknownTypeCommandSettings{}
				if err = unmarshal(&settings); err == nil {

					command.command = Command(&UnknownTypeCommand{})
					command.command.Settings(settings)

				}
			}
		}
	}

	return err
}
