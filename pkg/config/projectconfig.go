package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type server struct {
	Prefix string `yaml:"prefix"`
}

type projectConfig struct {
	Server server `yaml:"server"`
}

func loadProjectConfig() (*projectConfig, error) {
	file, err := filepath.Abs(filepath.Join("project", "project.yaml"))
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var projectConfigurations projectConfig
	if err := yaml.Unmarshal(data, &projectConfigurations); err != nil {
		return nil, err
	}

	return &projectConfigurations, nil
}
