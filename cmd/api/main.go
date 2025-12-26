package main

import (
	"github.com/JDarwind/go-skeleton-starter/pkg/config"
	"github.com/JDarwind/go-skeleton-starter/internals/routes"
	"github.com/JDarwind/go-skeleton-starter/pkg/server"
	"log"
	"net/http"
)



func main(){
	configurations := config.LoadConfig()

	mux := server.InitMuxWithRoutes( routes.NewRouter() )
	
	if err:= http.ListenAndServe(":" + configurations.ServerConfig.Port, mux); err != nil {
		log.Fatal(err)
	} 
}