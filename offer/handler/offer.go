package handler

import (
	"context"
	"errors"
	log "github.com/micro/go-micro/v2/logger"
	"offer/model"

	offer "offer/proto/offer"
)

type Offer struct{}

var OfferModel model.Offer

/**
设置offer
*/
func (e *Offer) SetOffer(ctx context.Context, request *offer.OfferRequest, response *offer.Response) error {
	if request.Name != "" && request.PostBackUrl != "" {

		if err := OfferModel.SetOffer(&model.Offer{
			Name:        request.Name,
			PostBackUrl: request.PostBackUrl,
		}); err != nil {
			return err
		}
		response.Msg = "设置成功"
		return nil
	} else {
		return errors.New("offer 参数不对，请重试设置")
	}
}

/**
获取offer
*/
func (e *Offer) GetOffer(ctx context.Context, request *offer.OfferRequest, response *offer.Response) error {
	if request.Id == "" {
		return errors.New("offer 参数不对，请重新设置")
	} else {
		log.Info("获取offer 成功 id", request.Id)
		OfferModel.GetOfferById(request.Id)
		response.Msg = "获取成功"
		return nil
	}
}
