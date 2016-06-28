package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	WUNDERTOOLS_CONFIG_APPLICATION_YAML_PATH = "settings.yml"
)

// iterate through all of the conf paths, and load application.yml
func (app *Application) from_ConfYaml() {

	for _, key := range app.Paths.OrderedConfPathKeys() {
		confPath, _ := app.Paths.Path(key)
		yamlFilePath := path.Join(confPath, WUNDERTOOLS_CONFIG_APPLICATION_YAML_PATH)

		if _, err := os.Stat(yamlFilePath); err == nil {
			yamlFile, err := ioutil.ReadFile(yamlFilePath)
			if err == nil {
				app.from_ConfYamlBytes(yamlFile)
			}
		}
	}

}

/**
 * configura an application from a yaml stream of Bytes
 * @TODO make this a reader?
 */
func (app *Application) from_ConfYamlBytes(yamlBytes []byte) {
	// parse the config file contents as a ConfSource_projectyaml object
	source := new(conf_Yaml)
	if err := yaml.Unmarshal(yamlBytes, source); err == nil {
		source.configureProject(app)
	}
}

/**
 * An Application configuration from Yaml
 *
 * This struct provides an interim format that can be
 * used by matching yml files, that need not exactly
 * match our Application object.
 */
type conf_Yaml struct {
	Project string `yaml:"Project,omitempty"`
	Author  string `yaml:"Author,omitempty"`

	Environment string `yaml:"Environment,omitempty"`

	Paths map[string]string `yaml:"Paths,omitempty"`

	Tokens map[string]string `yaml:"Tokens,omitempty"`

	Settings map[string]string `yaml:"Settings,omitempty"`
}

// Make a Yaml Conf apply configuration to a Application object
func (conf *conf_Yaml) configureProject(app *Application) {
	// set a project name

	if conf.Project != "" {
		app.Name = conf.Project
	}
	// set a author name
	if conf.Author != "" {
		app.Author = conf.Author
	}

	// set an environment string
	if conf.Environment != "" {
		app.Environment = conf.Environment
	}

	// set any paths
	for key, keyPath := range conf.Paths {
		app.Paths.SetPath(key, keyPath, false)
	}

}

// convert common boolean strings to actual boolean
func (conf *conf_Yaml) SettingStringToFlag(value string) bool {
	switch strings.ToLower(value) {
	case "y":
		fallthrough
	case "yes":
		fallthrough
	case "true":
		fallthrough
	case "1":
		return true

	default:
		return false
	}
}
