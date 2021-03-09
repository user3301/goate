package main

import (
	"context"
	"io/ioutil"
	"path/filepath"

	"github.com/user3301/goate/pkg/store"

	"github.com/user3301/goate/pkg/store/implementations"

	"github.com/user3301/goate/internal/validator"

	"gopkg.in/yaml.v2"
)

// Config main config for system
type Config struct {
	AppConfig        AppConfig        `yaml:"app" validate:"required"`
	PingServerConfig PingServerConfig `yaml:"ping" validate:"required"`
}

// AppConfig config for the app
type AppConfig struct {
	Port  int                     `yaml:"port" validate:"required"`
	Store *implementations.Config `yaml:"store"`
}

// PingServerConfig config for health check server
type PingServerConfig struct {
	Port int `yaml:"port" validate:"required"`
}

func (c *Config) Validate() error {
	return validator.Validate(c)
}

var defaultConfig = &Config{
	AppConfig: AppConfig{
		Port:  8080,
		Store: &implementations.Config{},
	},
	PingServerConfig: PingServerConfig{
		Port: 8082,
	},
}

func loadConfig(configFile string) (*Config, error) {
	if configFile == "" {
		return defaultConfig, nil
	}
	filePath := filepath.Clean(configFile)
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var config *Config
	if err := yaml.UnmarshalStrict(bytes, config); err != nil {
		return nil, err
	}
	return config, nil
}

func (c Config) GetStore(ctx context.Context) (store.UserStorer, error) {
	return c.AppConfig.Store.StoreFromConfig(ctx)
}
