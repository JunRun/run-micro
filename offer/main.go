package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	log "github.com/micro/go-micro/v2/logger"
	_ "github.com/micro/go-plugins/registry/etcd"
	"github.com/micro/go-plugins/registry/etcdv3"
	"offer/handler"
	"offer/model"
	_ "offer/model"
	offer "offer/proto/offer"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.offer"),
		micro.Version("latest"),
	)
	// Initialise service
	service.Init(micro.Registry(etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"run-micro_etcd_1:2379"}

	}))) // Register Handler
	offer.RegisterOfferHandler(service.Server(), new(handler.Offer))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.service.offer", service.Server(), new(subscriber.Offer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	defer model.CloseDB()
}
