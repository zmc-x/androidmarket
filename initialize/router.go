package initialize

import (
	"androidmarket/router"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	// 初始化一个route
	route := gin.Default()
	// router
	// 商城路由初始化
	mallroute := router.RouterApi.Mallrouter
	mallgroup := route.Group("/mall")
	{
		mallroute.Malluser.Initialize(mallgroup)
		mallroute.Address.Initialize(mallgroup)
		mallroute.MallGoods.Initialize(mallgroup)
		mallroute.MallCart.Initialize(mallgroup)
		mallroute.MallOrder.Initialize(mallgroup)
	}
	// 后台管理路由初始化
	manageroute := router.RouterApi.Managerouter
	managegroup := route.Group("/manage")
	{
		manageroute.Managegoods.Initialize(managegroup)
	}
	return route
}
