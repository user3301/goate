package main

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config main config for system
type Config struct {
	AppConfig AppConfig `yaml:"app" validate:"required"`
}

// AppConfig config for the app
type AppConfig struct {
	Port int `yaml:"port" validate:"required"`
}

var defaultConfig = &Config{AppConfig: AppConfig{Port: 8080}}

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
