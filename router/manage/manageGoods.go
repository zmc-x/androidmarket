package manage

import (
	v1 "androidmarket/api"
	"androidmarket/middleware"
	"github.com/gin-gonic/gin"
)

type Managegoods struct{}

func (m Managegoods) Initialize(router *gin.RouterGroup) {
	goods := router.Group("v1").Use(middleware.CheckToken(), middleware.Cors())
	{
		goods.POST("/goods/addmajorinfo", v1.Api.ManageGroup.AddGoodsMajorInfo)
		goods.POST("/goods/uploadimages/:goodsid", v1.Api.ManageGroup.AddGoodsImages)
		goods.POST("/goods/addspecifications", v1.Api.ManageGroup.AddGoodsSpecific)
		goods.POST("/goods/updateinfo", v1.Api.ManageGroup.UpdateGoodsinfo)
		goods.DELETE("/goods/deleteinfo", v1.Api.ManageGroup.DeleteGoodsinfo)
	}
}
