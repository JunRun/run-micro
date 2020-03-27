package main

import (
	"github.com/micro/go-micro"
	log "github.com/micro/go-micro/v2/logger"
	"offer/handler"
	"offer/model"
	_ "offer/model"
	offer "offer/proto/offer"
	"offer/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.offer"),
		micro.Version("latest"),
	)
	// Initialise service
	service.Init()
	// Register Handler
	offer.RegisterOfferHandler(service.Server(), new(handler.Offer))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.offer", service.Server(), new(subscriber.Offer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	defer model.CloseDB()
}
