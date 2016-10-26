package null

import (
	"github.com/james-nesbitt/wundertools-go/api/operation/command"
	"github.com/james-nesbitt/wundertools-go/api/operation/config"
)

type NullConfigWrapper struct{}

func (wrapper *NullConfigWrapper) Get(key string) (config.ConfigScopedValues, error) {
	return config.ConfigScopedValues{}, nil
}
func (wrapper *NullConfigWrapper) Set(key string, value config.ConfigScopedValues) error {
	return nil
}
func (wrapper *NullConfigWrapper) List(parent string) ([]string, error) {
	return []string{}, nil
}

type NullSettingWrapper struct{}

func (wrapper *NullSettingWrapper) Get(key string) (string, error) {
	return "", nil
}
func (wrapper *NullSettingWrapper) Set(key string, value string) error {
	return nil
}
func (wrapper *NullSettingWrapper) List(parent string) ([]string, error) {
	return []string{}, nil
}

type NullCommandWrapper struct{}

func (wrapper *NullCommandWrapper) Get(key string) (command.Command, error) {
	return nil, nil
}
func (wrapper *NullCommandWrapper) Set(key string, comm command.Command) error {
	return nil
}
func (wrapper *NullCommandWrapper) List(parent string) ([]string, error) {
	return []string{}, nil
}
