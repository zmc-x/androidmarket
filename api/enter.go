package api

import (
	"androidmarket/api/v1/mall"
	"androidmarket/api/v1/manage"
)

type ApiGroup struct {
	MallGroup   mall.MallGroup     // 商城前台api
	ManageGroup manage.ManageGroup // 商城后台api
}

var Api = new(ApiGroup)
