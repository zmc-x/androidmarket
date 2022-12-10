package mall

import (
	v1 "androidmarket/api"
	"androidmarket/middleware"
	"github.com/gin-gonic/gin"
)

type MallCart struct{}

func (receiver MallCart) Initialize(route *gin.RouterGroup) {
	cart := route.Group("v1").Use(middleware.CheckToken(), middleware.Cors())
	{
		cart.POST("/shoppingcart/add", v1.Api.MallGroup.Cartadd)
		cart.POST("/shoppingcart/updatecount", v1.Api.MallGroup.Cartupdate)
		cart.DELETE("/shoppingcart/delete", v1.Api.MallGroup.CartDelete)
		cart.GET("/shoppingcart/query", v1.Api.MallGroup.CartQuery)
		cart.POST("/shoppingcart/queryById", v1.Api.MallGroup.CartQueryById)
	}
}
