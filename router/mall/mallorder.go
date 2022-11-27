package mall

import (
	v1 "androidmarket/api"
	"androidmarket/middleware"
	"github.com/gin-gonic/gin"
)

type MallOrder struct{}

func (m MallOrder) Initialize(route *gin.RouterGroup) {
	order := route.Group("v1").Use(middleware.CheckToken())
	{
		order.POST("/order/create", v1.Api.MallGroup.OrderCreate)
		order.GET("/order/update/:updatetype", v1.Api.MallGroup.Orderupdate)
		order.GET("/order/query/:querytype", v1.Api.MallGroup.OrderQuery)
		order.GET("/order/queryinfo", v1.Api.MallGroup.OrderQueryInfo)
	}
}
