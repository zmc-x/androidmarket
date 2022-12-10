package mall

import (
	v1 "androidmarket/api"
	"androidmarket/middleware"
	"github.com/gin-gonic/gin"
)

type Address struct{}

func (m Address) Initialize(router *gin.RouterGroup) {
	address := router.Group("v1").Use(middleware.CheckToken())
	{
		address.POST("/address/add", v1.Api.MallGroup.Addaddress)
		address.POST("/address/update", v1.Api.MallGroup.Updateaddress)
		address.DELETE("/address/delete", v1.Api.MallGroup.Deleteaddresses)
		address.GET("/address/select", v1.Api.MallGroup.Selectaddresses)
	}
}
