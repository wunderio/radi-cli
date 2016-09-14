package bytesource

import (
	"errors"

	"gopkg.in/yaml.v2"
)

/**
 * Bytes stream handling based on a Command Set-Get-List handler which
 * will return bytes arrays for values, which will correspond to
 * the commands list
 */

/**
 * Command
 */

func NewYmlFileCommand_FromFileByteSource(file *FileByteSource) *YmlFileCommand {
	source := &YmlFileCommand{
		file:   file,
		values: CommandValues{},
	}

	source.Load()
	return source
}
func NewYmlFileCommand_FromFilePath(filePath string) *YmlFileCommand {
	return NewYmlFileCommand_FromFileByteSource(NewFileByteSource_FromPath(filePath))
}

//
type YmlFileCommand struct {
	file   *FileByteSource
	values CommandValues
	BaseByteArraySourceOperation
}

//
func (com *YmlFileCommand) Load() error {
	com.FromFile(com.file)
	com.values = CommandValues{}
	return yaml.Unmarshal(com.source, &com.values)
}

//
func (com *YmlFileCommand) Save() error {
	if source, err := yaml.Marshal(com.values); err == nil {
		com.source = source
		com.ToFile(com.file)
		return nil
	} else {
		return errors.New("Could not marshall config values to yml")
	}
}

//
func (com *YmlFileCommand) List(parent string) []string {
	keys := []string{}
	for key, _ := range com.values {
		keys = append(keys, key)
	}
	return keys
}

//
func (com *YmlFileCommand) Get(key string) (ymlCommand, bool) {
	value, found := com.values[key]
	return value, found
}

//
func (com *YmlFileCommand) Set(key string, value ymlCommand) bool {
	com.values[key] = value
	if err := com.Save(); err == nil {
		return true
	} else {
		return false
	}
}

type CommandValues map[string]ymlCommand

type ymlCommand interface{}
