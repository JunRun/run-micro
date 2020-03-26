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
	offerRequest := &offer.OfferRequest{}
	offerRequest.PostBackUrl = ctx.PostForm("postBackUrl")
	offerRequest.Name = ctx.PostForm("name")
	log.Info("设置offer", offerRequest.PostBackUrl, offerRequest.Name)
	res, err := ol.SetOffer(context.TODO(), offerRequest)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(200, res)
	}
}
