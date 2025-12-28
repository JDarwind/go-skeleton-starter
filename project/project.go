package project

import (
	"github.com/JDarwind/go-skeleton-starter/pkg/types"
)

func InitProject() (*types.ProjectConfig, error){
	projectConfig := createProjectConfig()
	return projectConfig, nil
}


func createProjectConfig() (*types.ProjectConfig) {
	return &types.ProjectConfig{
		Server: types.Server{
			Prefix: "/",
		},
	}
}