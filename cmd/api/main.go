package main

import (
	"log"
	"net/http"

	"github.com/JDarwind/go-skeleton-starter/internals/routes"
	"github.com/JDarwind/go-skeleton-starter/pkg/config"
	"github.com/JDarwind/go-skeleton-starter/pkg/server"
)

func main() {
	configurations := config.NewConfigManager(nil).GetConfig()

	mux := server.InitMuxWithRoutes(routes.NewRouter())

	if err := http.ListenAndServe(":"+configurations.ProjectConfig.Server.Port, mux); err != nil {
		log.Fatal(err)
	}
}
