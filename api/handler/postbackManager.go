/**
 *
 * @Description:  回传模块 管理
 * @Version: 1.0.0
 * @Date: 2020-03-26 10:04
 */
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
)

type PostBackManager struct {
}

func (p *PostBackManager) SetPostBack(ctx *gin.Context) {
	logger.Info("设置postback 成功")
	ctx.JSON(200, "设置成功")
}
