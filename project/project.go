package project

import (
	"github.com/JDarwind/go-skeleton-starter/pkg/types"
	"os"
)

func getEnv( name string, defaultValue string ) ( string ) {
	env:= os.Getenv( name )

	if env != "" {
		return env
	}
	
	return defaultValue
}

func InitProject() *types.ProjectConfig{
	return &types.ProjectConfig{
		Server: types.Server{
			Prefix: "/",
			Port: getEnv("APP_PORT", "8081"),
		},
	}
}