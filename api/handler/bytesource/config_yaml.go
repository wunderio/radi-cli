package bytesource

import (
	"errors"

	"gopkg.in/yaml.v2"

	"github.com/james-nesbitt/wundertools-go/api/operation/config"
)

/**
 * Bytes stream handling based on a Config Set-Get handler which
 * will return bytes arrays for values, which will correspond to
 * the configuration of other operations.
 */

/**
 * Config
 */

func NewYmlFileConfig_FromFileByteSource(file *FileByteSource) *YmlFileConfig {
	source := &YmlFileConfig{
		file:   file,
		values: ConfigValues{},
	}

	source.Load()
	return source
}
func NewYmlFileConfig_FromFilePath(filePath string) *YmlFileConfig {
	return NewYmlFileConfig_FromFileByteSource(NewFileByteSource_FromPath(filePath))
}

//
type YmlFileConfig struct {
	file   *FileByteSource
	values ConfigValues
	BaseByteArraySourceOperation
}

//
func (config *YmlFileConfig) Load() error {
	config.FromFile(config.file)
	config.values = ConfigValues{}
	return yaml.Unmarshal(config.source, &config.values)
}

//
func (config *YmlFileConfig) Save() error {
	if source, err := yaml.Marshal(config.values); err == nil {
		config.source = source
		config.ToFile(config.file)
		return nil
	} else {
		return errors.New("Could not marshall config values to yml")
	}
}

//
func (config *YmlFileConfig) List(parent string) []string {
	keys := []string{}
	for key, _ := range config.values {
		keys = append(keys, key)
	}
	return keys
}

//
func (config *YmlFileConfig) Get(key string) (string, bool) {
	value, found := config.values[key]
	return value, found
}

//
func (config *YmlFileConfig) Set(key string, value string) bool {
	config.values[key] = value
	if err := config.Save(); err == nil {
		return true
	} else {
		return false
	}
}

type ConfigValues map[string]string

//
type YmlFileConfigListOperation struct {
	config.BaseConfigListOperation
	source YmlFileConfig
}

//
type YmlFileConfigGetOperation struct {
	config.BaseConfigGetOperation
	source YmlFileConfig
}

//
type YmlFileConfigSetOperation struct {
	config.BaseConfigSetOperation
	source YmlFileConfig
}
