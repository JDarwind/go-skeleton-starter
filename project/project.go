package project

import (
	"github.com/JDarwind/go-skeleton-starter/pkg/types"
	"os"
)

func getEnv(name string, defaultValue string) string {
	env := os.Getenv(name)

	if env != "" {
		return env
	}

	return defaultValue
}

func InitProject() *types.ProjectConfig {
	return &types.ProjectConfig{
		Server: types.Server{
			Prefix: "/",
			Port:   getEnv("APP_PORT", "8081"),
		},
	}
}
//Default No operation on application config, you cna overrid it here for global configuation
func InitApplicationConfig(applicationConfig any) any{
	if applicationConfig != nil{
		return applicationConfig
	}
	return nil
}