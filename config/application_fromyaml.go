package config

import (
	"path"
	"os"
	"ioutil"
	"gopkg.in/yaml.v2"
)

const (
	WUNDERTOOLS_CONFIG_APPLICATION_YAML_PATH = "application.yml"
)

// iterate through all of the conf paths, and load application.yml
func (app *Application) from_ConfYaml(paths *Paths) {

	for _, key := range paths.OrderedConfPathKeys() {
		path, _ := paths.Path(key)
		yamlFilePath := path.Join(path, WUNDERTOOLS_CONFIG_APPLICATION_YAML_PATH)

		if _, err := os.Stat(yamlFilePath); err == nil {
			yamlFile, err := ioutil.ReadFile(yamlFilePath)
			if err == nil {	
				project.from_ConfYamlBytes(logger.MakeChild(yamlFilePath), yamlFile)
			}
		}
	}

}
/**
 * configura an application from a yaml stream of Bytes
 * @TODO make this a reader?
 */
func (app *Application) from_ConfYamlBytes(yamlBytes []byte) bool {
	// parse the config file contents as a ConfSource_projectyaml object
	source := new(conf_Yaml)
	if err := yaml.Unmarshal(yamlBytes, source); err != nil {
		logger.Warning("YAML parsing error : " + err.Error())
		return false
	}
	logger.Debug(log.VERBOSITY_DEBUG_STAAAP, "YAML source:", *source)

	return source.configureProject(logger, app)
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
func (conf *conf_Yaml) configureProject(app *Application) bool {
	// set a project name

	if conf.Project != "" {
		project.Name = conf.Project
	}
	// set a author name
	if conf.Author != "" {
		project.Author = conf.Author
	}

	// set an environment string
	if conf.Environment != "" {
		project.Environment = conf.Environment
	}

	// set any paths
	for key, keyPath := range conf.Paths {
		project.Paths.SetPath(key, keyPath, true)
	}

	// set any tokens
	for key, value := range conf.Tokens {
		project.Tokens.SetToken(key, value)
	}

	/**
	 * Yaml Settings set Project Flags
	 */
	for key, value := range conf.Settings {
		switch key {
		case "UsePathsAsTokens":
			project.UsePathsAsTokens = conf.SettingStringToFlag(value)
		case "UseEnvVariablesAsTokens":
			project.UseEnvVariablesAsTokens = conf.SettingStringToFlag(value)
		}
	}

	logger.Debug(log.VERBOSITY_DEBUG_LOTS, "Configured project from YAML conf", project)
	return true
}

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



