package config

import (
	"log"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
)

type Config struct{
	ProjectConfig projectConfig
	
	ServerConfig struct {
		Port string
	}
	ApplicationConfigs any
}

type ConfigOptions struct {
	EnvironmentFileLocaltion string
	ApplicationConfigs any
}


type ConfigManager struct {
	environmentFileLocaltion string
	applicationConfigs any
	configurations *Config
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

func (cm *ConfigManager) loadConfig() *Config {
    if cm.configurations != nil {
        return cm.configurations
    }

    cm.configurations = &Config{}

    envFile, err := filepath.Abs(cm.environmentFileLocaltion)
    if err != nil {
        log.Fatal(err)
    }

    _ = godotenv.Load(envFile)

    cm.initConfiguartionObject()

    cm.configurations.ApplicationConfigs = cm.applicationConfigs
    return cm.configurations
}

func GetConfigManager() (*ConfigManager){
	if configrationManager == nil{
		panic("Config Manager not initialized")
	}
	return configrationManager
}

func(configManager *ConfigManager) GetConfig()  ( *Config ){
	return configManager.configurations
}


func (configManager *ConfigManager) initConfiguartionObject(){
	projectConfig, err := loadProjectConfig()
	
	if err != nil{
		log.Fatal(err)
	}
	
	configManager.configurations.ServerConfig.Port = configManager.getEnv("APP_PORT", "")
	configManager.configurations.ProjectConfig.Server.Prefix = projectConfig.Server.Prefix
}


func (configManager *ConfigManager) getEnv( name string, defaultValue string ) ( string ) {
	env:= os.Getenv( name )

	if env != "" {
		return env
	}
	
	return defaultValue
}