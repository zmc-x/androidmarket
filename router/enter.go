package router

import (
	"androidmarket/router/mall"
	"androidmarket/router/manage"
)

type Router struct {
	Mallrouter   mall.MallRouteGroup     // 商城前台路由管理
	Managerouter manage.ManageRouteGroup // 商城后台路由管理
}

var RouterApi = new(Router)
