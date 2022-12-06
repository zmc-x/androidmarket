package mall

import (
	"androidmarket/middleware"
	mallresp "androidmarket/model/common/response"
	"androidmarket/model/mall/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Mallorder struct{}

// OrderCreate 生成订单
func (m Mallorder) OrderCreate(c *gin.Context) {
	input := request.Createoder{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		mallresp.FailParam(c)
		return
	}
	if reserr, status, msg := OrderCreate(input, middleware.Uid); status {
		mallresp.OkWithDetail(c, nil, msg)
	} else {
		mallresp.FailWithDetail(c, map[string]interface{}{"error": reserr.Error()}, msg)
	}
}

// Orderupdate 修改订单的状态
func (m Mallorder) Orderupdate(c *gin.Context) {
	// 获取参数
	updatetype, mid, uid := c.Param("updatetype"), c.Query("orderid"), middleware.Uid
	orderid, _ := strconv.Atoi(mid)
	if msg, status := OrderUpdate(updatetype, orderid, uid); status {
		mallresp.OkWithDetail(c, nil, msg)
	} else {
		mallresp.FailParam(c)
	}
}

// OrderQuery 查询订单
func (m Mallorder) OrderQuery(c *gin.Context) {
	// 获取参数
	querytype, uid := c.Param("querytype"), middleware.Uid
	if res, status, msg := OrderQuery(querytype, uid); status {
		mallresp.OkWithDetail(c, res, msg)
	} else {
		mallresp.FailParam(c)
	}
}

// OrderQueryInfo 查询订单详细信息
func (m Mallorder) OrderQueryInfo(c *gin.Context) {
	// 获取参数
	mid, uid := c.Query("orderid"), middleware.Uid
	orderid, _ := strconv.Atoi(mid)
	res := OrderQueryInfo(uid, orderid)
	mallresp.OkWithDetail(c, res, "")
}
