package mall

import (
	mallresp "androidmarket/model/common/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MallGoods struct{}

// Showgoodsinfo 查询商品的单个信息
func (m MallGoods) Showgoodsinfo(c *gin.Context) {
	goodsid, _ := strconv.Atoi(c.Query("goodsid"))
	specificationid, _ := strconv.Atoi(c.Query("specificationid"))
	if status, info := ShowGoodsInfo(goodsid, specificationid); status {
		mallresp.OkWithDetail(c, info, "查询成功！")
	} else {
		mallresp.FailWithDetail(c, nil, "查询失败！")
	}
}

// QueryGoodsByType 通过商品类别来查询相关信息
func (m MallGoods) QueryGoodsByType(c *gin.Context) {
	goodstype := c.Query("goodstype")
	if status, res := QueryGoodsByType(goodstype); status {
		mallresp.OkWithDetail(c, res, "查询成功！")
	} else {
		mallresp.FailWithDetail(c, nil, "查询失败！")
	}
}

// GoodsHomeInfo 首页商品信息
func (m MallGoods) GoodsHomeInfo(c *gin.Context) {
	res := GoodsHomeInfo()
	mallresp.OkWithDetail(c, res, "查询成功！")
}

// GoodsInfo 查询商品在订单中的相关信息
func (m MallGoods) GoodsInfo(c *gin.Context) {
	goodsid, _ := strconv.Atoi(c.Query("goodsid"))
	specificationid, _ := strconv.Atoi(c.Query("specificationid"))
	res := Goodsinfo(goodsid, specificationid)
	mallresp.OkWithDetail(c, res, "查询成功！")
}

// QueryByName 通过商品名称来查询相关信息
func (m MallGoods) QueryByName(c *gin.Context) {
	goodsname := c.Query("goodsname")
	if res, msg, status := QueryByName(goodsname); status {
		mallresp.OkWithDetail(c, res, msg)
	} else {
		mallresp.FailWithDetail(c, res, msg)
	}
}
