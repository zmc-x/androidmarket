package mall

import (
	"androidmarket/middleware"
	mallresp "androidmarket/model/common/response"
	"androidmarket/model/mall/request"
	"github.com/gin-gonic/gin"
)

type MallCart struct{}

// Cartadd 添加商品到购物车中
func (m MallCart) Cartadd(c *gin.Context) {
	input := request.AddCart{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		mallresp.FailParam(c)
		return
	}
	if reserr, status, msg := CartAdd(input, middleware.Uid); status {
		mallresp.OkWithDetail(c, nil, msg)
	} else {
		if reserr != nil {
			mallresp.FailWithDetail(c, map[string]interface{}{"error": reserr.Error()}, msg)
		} else {
			mallresp.FailWithDetail(c, nil, msg)
		}
	}
}

// Cartupdate 修改商品在购物车中的数量
func (m MallCart) Cartupdate(c *gin.Context) {
	input := request.UpdateCount{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		mallresp.FailParam(c)
		return
	}
	if msg, newCount, status := Cartupdate(input, middleware.Uid); status {
		mallresp.OkWithDetail(c, map[string]interface{}{"newcount": newCount}, msg)
	} else {
		mallresp.FailWithDetail(c, nil, msg)
	}
}

// CartDelete 删除购物车中商品
func (m MallCart) CartDelete(c *gin.Context) {
	input := request.Cartdelete{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		mallresp.FailParam(c)
		return
	}
	mallresp.OkWithDetail(c, nil, Cartdelete(input, middleware.Uid))
}

// CartQuery 查询购物车中商品信息
func (m MallCart) CartQuery(c *gin.Context) {
	res := CartQuery(middleware.Uid)
	mallresp.OkWithDetail(c, res, "")
}

// CartQueryById 查询购物车中商品信息通过id
func (m MallCart) CartQueryById(c *gin.Context) {
	queryparam := request.CartQueryById{}
	paramerr := c.ShouldBindJSON(&queryparam)
	if paramerr != nil {
		mallresp.FailParam(c)
	}
	res := CartQueryById(middleware.Uid, queryparam.Cartids)
	mallresp.OkWithDetail(c, res, "")
}
