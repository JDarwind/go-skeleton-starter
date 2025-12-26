package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Prefix string `yaml:"prefix"`
}

type ProjectConfig struct {
	Server Server `yaml:"server"`
}

func loadProjectConfig() (*ProjectConfig, error) {
	file, err := filepath.Abs(filepath.Join("project", "project.yaml"))
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var projectConfigurations ProjectConfig
	if err := yaml.Unmarshal(data, &projectConfigurations); err != nil {
		return nil, err
	}

	return &projectConfigurations, nil
}
