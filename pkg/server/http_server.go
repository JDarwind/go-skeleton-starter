package server

import (
	"github.com/JDarwind/go-skeleton-starter/pkg/config"
	"net/http"
	"strings"
)

func InitMuxWithRoutes(router *http.ServeMux) *http.ServeMux {
	cfg := config.LoadConfig()

	server := http.NewServeMux()

	rawPrefix := strings.TrimSpace(cfg.ProjectConfig.Server.Prefix)
	if rawPrefix == "" || rawPrefix == "/" {
		server.Handle("/", router)
		return server
	}

	prefix := strings.TrimRight(rawPrefix, "/")

	server.Handle(prefix+"/", http.StripPrefix(prefix, router))

	server.Handle(prefix, http.RedirectHandler(prefix +"/", http.StatusMovedPermanently))

	return server
}