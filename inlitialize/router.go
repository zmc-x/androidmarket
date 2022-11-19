package inlitialize

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
		mallroute.Malluser.Inlitialize(mallgroup)
		mallroute.Address.Inlitilize(mallgroup)
		mallroute.MallGoods.Inlitialize(mallgroup)
		mallroute.MallCart.Inlitialize(mallgroup)
		mallroute.MallOrder.Inlitialize(mallgroup)
	}
	// 后台管理路由初始化
	manageroute := router.RouterApi.Managerouter
	managegroup := route.Group("/manage")
	{
		manageroute.Managegoods.Inlitialize(managegroup)
	}
	return route
}
