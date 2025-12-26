package config

import (
	"log"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
)

type Config struct{
	ProjectConfig ProjectConfig
	
	ServerConfig struct {
		Port string
	}
}

var configurations *Config = nil;

func LoadConfig () ( *Config ){
	
	if configurations != nil {
		return configurations
	}
	
	configurations = &Config{}
	
	envFile, err:= filepath.Abs(".env")
	
	if err != nil{
		log.Fatal(err)
	}
	
	_= godotenv.Load(envFile)
	
	initConfiguartionObject()
	
	return configurations
}

func initConfiguartionObject(){
	projectConfig, err := loadProjectConfig()
	
	if err != nil{
		log.Fatal(err)
	}
	
	configurations.ServerConfig.Port = getEnv("APP_PORT", "")
	configurations.ProjectConfig.Server.Prefix = projectConfig.Server.Prefix
}


func getEnv( name string, defaultValue string ) ( string ) {
	env:= os.Getenv( name )

	if env != "" {
		return env
	}
	
	return defaultValue
}