package service

import (
	"androidmarket/service/mall"
	"androidmarket/service/manage"
)

type ServiceGroup struct {
	MallGroup   mall.MallServiceGroup // mall前台服务
	ManageGroup manage.ManageGroup    // manage后台服务
}

var Service = new(ServiceGroup)
