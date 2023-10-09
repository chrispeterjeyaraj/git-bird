package config

import (
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Version string `yaml:"version"`
}

func LoadConfigFromBytes(data []byte) (AppConfig, error) {
	var config AppConfig

	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}
