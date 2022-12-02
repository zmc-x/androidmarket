package mall

import (
	"androidmarket/global"
	"androidmarket/model/mall"
	"androidmarket/model/mall/request"
	"androidmarket/model/mall/response"
	"fmt"
	"time"
)

type MallOrder struct{}

// CreateOrder 生成订单
func (m MallOrder) CreateOrder(param request.Createoder, uid string) (error, bool, string) {
	now := time.Now()
	// 开启事务
	tx := global.GlobalDB.Begin()
	// 查看是否为从购物车中生成订单
	if len(param.Cartids) != 0 {
		// 删除购物车中的商品
		tx.Delete(&mall.Cart{}, param.Cartids)
	}
	// 首先操纵order表
	res := tx.Create(&mall.Order{
		Uid:        uid,
		AddressId:  param.Order.Addressid,
		Allprice:   param.Order.Allprice,
		Createdat:  fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()),
		Finishedat: fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()),
		Status:     1,
	})
	if res.Error != nil {
		tx.Rollback()
		return res.Error, false, "生成订单失败！"
	}
	// 获取orderid
	temp := mall.Order{}
	tx.Last(&temp)
	// 操作orderitem表
	for _, v := range param.Orderitem {
		res := tx.Create(&mall.Orderitem{
			OrderId:    temp.Id,
			Color:      v.Color,
			Specific:   v.Specific,
			Count:      v.Count,
			Price:      v.Price,
			GoodsName:  v.Goodsname,
			CoverImage: v.Coverimage,
		})
		if res.Error != nil {
			tx.Rollback()
			return res.Error, false, "生成订单失败！"
		}
	}
	tx.Commit()
	return nil, true, "生成订单成功！"
}

// Updateorder 修改订单的状态
func (m MallOrder) Updateorder(updatetype string, orderid int, uid string) (string, bool) {
	now := time.Now()
	if updatetype == "cancel" {
		global.GlobalDB.Model(&mall.Order{}).Where("uid = ? and id = ?", uid, orderid).Select("status", "finishedat").Updates(mall.Order{
			Status:     2,
			Finishedat: fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()),
		})
		return "取消订单成功！", true
	} else if updatetype == "complete" {
		global.GlobalDB.Model(&mall.Order{}).Where("uid = ? and id = ?", uid, orderid).Select("status", "finishedat").Updates(mall.Order{
			Status:     3,
			Finishedat: fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()),
		})
		return "订单完成！", true
	} else {
		return "", false
	}
}

// OrderQuery 订单查询
func (m MallOrder) OrderQuery(querytype string, uid string) ([]response.Orders, bool, string) {
	mid := make([]mall.Order, 0)
	if querytype == "cancel" {
		global.GlobalDB.Where("uid = ? and status = 2", uid).Find(&mid)
	} else if querytype == "complete" {
		global.GlobalDB.Where("uid = ? and status = 3", uid).Find(&mid)
	} else if querytype == "all" {
		global.GlobalDB.Where("uid = ? ", uid).Find(&mid)
	} else if querytype == "unfinished" {
		global.GlobalDB.Where("uid = ? and status = 1", uid).Find(&mid)
	} else {
		return nil, false, ""
	}
	// 对查询到的id进行处理
	res := make([]response.Orders, 0)
	for _, v := range mid {
		temp := make([]response.Goodsinfo, 0)
		// 查询每个订单中的商品信息
		global.GlobalDB.Model(&mall.Orderitem{}).Where("order_id = ?", v.Id).Find(&temp)
		res = append(res, response.Orders{
			Orderid:  v.Id,
			Allprice: v.Allprice,
			Goods:    temp,
		})
	}
	return res, true, "查询成功！"
}

// OrderQueryInfo 查询订单详细信息
func (m MallOrder) OrderQueryInfo(uid string, orderid int) response.Orderinfo {
	mid := struct {
		Province       string    `json:"province"`       // 省
		City           string    `json:"city"`           // 市
		County         string    `json:"county"`         // 区 / 县
		Detaillocation string    `json:"detaillocation"` // 详细地址
		Allprice       float64   `json:"allprice"`       // 总价格
		Createdat      time.Time `json:"createdat"`      // 创建时间
		Finishedat     time.Time `json:"finishedat"`     // 取消/完成时间
		Status         int       `json:"status"`         // 订单状态
	}{}
	global.GlobalDB.Model(&mall.Order{}).Where("orders.uid = ? and orders.id = ?", uid, orderid).Select("orders.allprice, orders.status, orders.createdat, orders.finishedat, addresses.province, addresses.city, addresses.county, addresses.detaillocation").Joins("join addresses on orders.address_id = addresses.id").Scan(&mid)
	// 将mid赋值给res
	res := response.Orderinfo{
		Province:       mid.Province,
		County:         mid.County,
		City:           mid.City,
		Detaillocation: mid.Detaillocation,
		Allprice:       mid.Allprice,
		Createdat:      fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", mid.Createdat.Year(), mid.Createdat.Month(), mid.Createdat.Day(), mid.Createdat.Hour(), mid.Createdat.Minute(), mid.Createdat.Second()),
		Finishedat:     fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", mid.Finishedat.Year(), mid.Finishedat.Month(), mid.Finishedat.Day(), mid.Finishedat.Hour(), mid.Finishedat.Minute(), mid.Finishedat.Second()),
		Status:         mid.Status,
	}
	temp := make([]mall.Orderitem, 0)
	global.GlobalDB.Where("order_id = ?", orderid).Find(&temp)
	for _, v := range temp {
		res.Goodsinfo = append(res.Goodsinfo, response.Goodsinfo{
			Specific:   v.Specific,
			GoodsName:  v.GoodsName,
			Count:      v.Count,
			Color:      v.Color,
			CoverImage: "https://cdn.zmcicloud.cn/" + v.CoverImage,
			Price:      v.Price,
		})
	}
	return res
}
