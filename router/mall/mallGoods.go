package mall

import (
	v1 "androidmarket/api"
	"androidmarket/middleware"
	"github.com/gin-gonic/gin"
)

type MallGoods struct{}

func (g MallGoods) Initialize(router *gin.RouterGroup) {
	goods := router.Group("v1").Use(middleware.CheckToken(), middleware.Cors())
	{
		goods.GET("/goods/showgoodsinfo", v1.Api.MallGroup.Showgoodsinfo)
		goods.GET("/goods/Bytype", v1.Api.MallGroup.QueryGoodsByType)
		goods.GET("/goods/queryById", v1.Api.MallGroup.GoodsInfo)
	}
	// 此处无需鉴权
	goodsTwo := router.Group("v1").Use(middleware.Cors())
	{
		goodsTwo.GET("/goods/marketinfo", v1.Api.MallGroup.GoodsHomeInfo)
	}
}
