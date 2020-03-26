package main

import (
	"api/router"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/web"
)

func main() {
	// New Service
	service := web.NewService(web.Name("go.micro.api.api"))

	// Initialise service
	service.Init()
	gin := router.ApiService()
	service.Handle("/", gin)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
