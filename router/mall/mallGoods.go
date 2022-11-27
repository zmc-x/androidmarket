package mall

import (
	v1 "androidmarket/api"
	"androidmarket/middleware"
	"github.com/gin-gonic/gin"
)

type MallGoods struct{}

func (g MallGoods) Initialize(router *gin.RouterGroup) {
	goods := router.Group("v1").Use(middleware.CheckToken())
	{
		goods.GET("/goods/showgoodsinfo", v1.Api.MallGroup.Showgoodsinfo)
		goods.GET("/goods/Bytype", v1.Api.MallGroup.QueryGoodsByType)
	}
}
