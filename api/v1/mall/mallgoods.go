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
	if status, info := ShowGoodsInfo(goodsid); status {
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
