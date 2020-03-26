/**
 *
 * @Description:  offer 管理
 * @Version: 1.0.0
 * @Date: 2020-03-26 13:22
 */
package handler

import (
	offer "api/proto/offer"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	log "github.com/micro/go-micro/v2/logger"
)

type OfferManager struct {
}

var ol = offer.NewOfferService("", client.DefaultClient)

func (o *OfferManager) SetOffer(ctx *gin.Context) {
	var offer = &offer.OfferRequest{}
	if err := ctx.Bind(offer); err != nil {
		log.Error("设置offer 出错", err)
	}
	res, err := ol.SetOffer(context.TODO(), offer)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(200, res)
	}
}
