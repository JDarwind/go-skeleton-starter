package config

import (
	"log"
	"path/filepath"

	"github.com/JDarwind/go-skeleton-starter/pkg/types"
	"github.com/JDarwind/go-skeleton-starter/project"
	"github.com/joho/godotenv"
)

type ConfigOptions struct {
	EnvironmentFileLocaltion string
	ApplicationConfigs any
}


type ConfigManager struct {
	environmentFileLocaltion string
	applicationConfigs any
	configurations *types.Config
}


var configrationManager *ConfigManager = nil


func NewConfigManager(options *ConfigOptions) *ConfigManager {
	
	if options == nil{
		options = &ConfigOptions{
			EnvironmentFileLocaltion: "",
			ApplicationConfigs: nil,
		}
	}
	
    envFile := ".env"
    if options != nil && options.EnvironmentFileLocaltion != "" {
        envFile = options.EnvironmentFileLocaltion
    }

    var appConfig any
    if options != nil && options.ApplicationConfigs != nil {
        appConfig = options.ApplicationConfigs
    }

    cm := &ConfigManager{
        environmentFileLocaltion: envFile,
        applicationConfigs:       appConfig,
        configurations:           nil,
    }

    cm.loadConfig()
    
    configrationManager = cm
    return cm
}

func (cm *ConfigManager) loadConfig() *types.Config {
    if cm.configurations != nil {
        return cm.configurations
    }

    cm.configurations = &types.Config{}

    envFile, err := filepath.Abs(cm.environmentFileLocaltion)
    if err != nil {
        log.Fatal(err)
    }

    _ = godotenv.Load(envFile)

    cm.configurations.ProjectConfig = project.InitProject()
    cm.configurations.ApplicationConfigs = cm.applicationConfigs
    return cm.configurations
}

func GetConfigManager() (*ConfigManager){
	if configrationManager == nil{
		panic("Config Manager not initialized")
	}
	return configrationManager
}

func(configManager *ConfigManager) GetConfig()  ( *types.Config ){
	return configManager.configurations
}