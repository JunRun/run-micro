package main

import (
	"api/router"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	// New Service
	service := micro.NewService(micro.Name("go.micro.api.api"))

	// Initialise service
	service.Init(micro.Registry(etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"run-micro_etcd_1:2379"}
	})))
	gin := router.ApiService()
	// Run service
	if err := gin.Run(":8080"); err != nil {
		logger.Error(err)
	}

}
