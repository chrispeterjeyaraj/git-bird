package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Version string `yaml:"version"`
}

func LoadConfig(filename string) (AppConfig, error) {
	var config AppConfig

	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	if err := yaml.Unmarshal(configFile, &config); err != nil {
		return config, err
	}

	return config, nil
}
