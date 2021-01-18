package bot

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	Discord struct {
		ClientId     string `yaml:"id"`
		ClientSecret string `yaml:"secret"`
		Token        string `yaml:"token"`
	}
	Postgres struct {
		HostAddr         string `yaml:"addr"`
		HostPort         int    `yaml:"port"`
		DatabaseUser     string `yaml:"user"`
		DatabasePassword string `yaml:"password"`
		DatabaseName     string `yaml:"database"`
	}
}

func NewConfig(yamlConfig []byte) (*Config, error) {
	config := Config{}
	err := yaml.Unmarshal(yamlConfig, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
