package mall

import (
	"androidmarket/middleware"
	mallresp "androidmarket/model/common/response"
	"androidmarket/model/mall/request"
	"github.com/gin-gonic/gin"
)

type Malladdress struct{}

// Addaddress 添加地址
func (m Malladdress) Addaddress(c *gin.Context) {
	addinfo := request.Address{}
	err := c.ShouldBindJSON(&addinfo)
	if err != nil {
		mallresp.FailParam(c)
		return
	}
	if msg, status, err := Addaddress(addinfo, middleware.Uid); status {
		mallresp.OkWithDetail(c, nil, msg)
	} else {
		if err != nil {
			mallresp.FailWithDetail(c, map[string]interface{}{"error": err.Error()}, msg)
		} else {
			mallresp.FailWithDetail(c, nil, msg)
		}
	}
}

// Updateaddress 修改地址
func (m Malladdress) Updateaddress(c *gin.Context) {
	updateinfo := request.Address{}
	err := c.ShouldBindJSON(&updateinfo)
	if err != nil {
		mallresp.FailParam(c)
		return
	}
	if msg, status, err := Updateaddress(updateinfo, middleware.Uid); status {
		mallresp.OkWithDetail(c, nil, msg)
	} else {
		if err != nil {
			mallresp.FailWithDetail(c, map[string]interface{}{"error": err.Error()}, msg)
		} else {
			mallresp.FailWithDetail(c, nil, msg)
		}
	}
}

// Deleteaddresses 删除地址
func (m Malladdress) Deleteaddresses(c *gin.Context) {
	requestid := request.Addressesid{}
	err := c.ShouldBindJSON(&requestid)
	if err != nil {
		mallresp.FailParam(c)
		return
	}
	if msg, status := Deleteaddresses(requestid.Ids, middleware.Uid); status {
		mallresp.OkWithDetail(c, nil, msg)
	} else {
		mallresp.FailWithDetail(c, nil, msg)
	}
}

// Selectaddresses 查询地址
func (m Malladdress) Selectaddresses(c *gin.Context) {
	addresses := Selectaddresses(middleware.Uid)
	mallresp.OkWithDetail(c, addresses, "查询成功！")
}
