package mall

import (
	v1 "androidmarket/api"
	"androidmarket/middleware"
	"github.com/gin-gonic/gin"
)

type Malluser struct{}

func (m Malluser) Inlitialize(router *gin.RouterGroup) {
	usergroup := router.Group("v1")
	{
		usergroup.POST("/user/Login", v1.Api.MallGroup.Userlogin)
		usergroup.POST("/user/signup", v1.Api.MallGroup.Usersignup)
	}
	usergrouptwo := router.Group("v1").Use(middleware.CheckToken())
	{
		usergrouptwo.POST("/user/updatepass", v1.Api.MallGroup.Userupdatepass)
	}
}
