/**
 *
 * @Description: 上游路由控制
 * @Version: 1.0.0
 * @Date: 2020-03-26 09:47
 */
package router

import (
	"api/handler"
	"github.com/gin-gonic/gin"
)

func ApiService() *gin.Engine {
	router := gin.Default()

	//所有请求都以  /api/* 拦截
	api := router.Group("/api")

	// 回传模块 管理
	postBack := new(handler.PostBackManager)
	postBackRouter := api.Group("/postBack")
	postBackRouter.POST("/add", postBack.SetPostBack)

	offer := new(handler.OfferManager)
	offerRouter := api.Group("/offer")
	offerRouter.POST("/set", offer.SetOffer)

	return router
}
