package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	offer "offer/proto/offer"
)

type Offer struct{}

func (e *Offer) Handle(ctx context.Context, msg *offer.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *offer.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
