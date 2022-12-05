package mall

import (
	"androidmarket/global"
	"androidmarket/model/mall"
	"androidmarket/model/mall/request"
	"androidmarket/model/mall/response"
)

type MallCart struct{}

// Cartadd 添加商品到购物车
func (m MallCart) Cartadd(param request.AddCart, uid string) (error, bool, string) {
	// 判断相关数据是否合法
	temp, tempspecific := mall.Cart{}, mall.Specification{}
	midcart := global.GlobalDB.Where("specification_id = ? and goods_id = ?", param.Specificationid, param.Goodsid).Find(&temp)
	// 查询对应商品规格的数量
	global.GlobalDB.Where("specification_id = ? and goods_id = ?", param.Specificationid, param.Goodsid).Find(&tempspecific)
	// 是否存在数据
	// 若存在数据则返回提示信息
	// 反之则进行添加的操作
	if midcart.RowsAffected != 0 {
		return nil, false, "该商品已经存在于购物车中，无需在进行重复的添加！"
	} else {
		mid := global.GlobalDB.Where("specification_id = ? and count >= ?", param.Specificationid, param.Count).Find(&mall.Specification{})
		if mid.RowsAffected == 0 {
			return nil, false, "数据非法！"
		}
		res := global.GlobalDB.Create(&mall.Cart{
			GoodsId:         param.Goodsid,
			SpecificationId: param.Specificationid,
			Count:           param.Count,
			Uid:             uid,
		})
		if res.Error != nil {
			return res.Error, false, "添加失败！"
		} else {
			return nil, true, "添加成功！"
		}
	}
	return nil, true, "添加成功！"
}

// Cartupdatecount 更新购物车中的数量
func (m MallCart) Cartupdatecount(param request.UpdateCount, uid string) (string, int, bool) {
	Ctemp, Stemp := mall.Cart{}, mall.Specification{}
	res := global.GlobalDB.Where("id = ? and uid = ? ", param.Cartid, uid).Find(&Ctemp)
	if res.RowsAffected == 0 {
		return "非法修改", 0, false
	}
	global.GlobalDB.Where("specification_id = ?", param.Specificationid).Find(&Stemp)
	if param.Newcount > Stemp.Count {
		global.GlobalDB.Model(&Ctemp).Update("count", Stemp.Count)
		return "Sorry，该商品没有这么多的库存！", Stemp.Count, true
	} else {
		global.GlobalDB.Model(&Ctemp).Update("count", param.Newcount)
		return "", param.Newcount, true
	}
}

// Cartdelete 删除购物车中的相关的商品
func (m MallCart) Cartdelete(param request.Cartdelete, uid string) string {
	// 循环遍历
	for _, v := range param.Deletegoods {
		global.GlobalDB.Where("id = ? and uid = ?", v, uid).Delete(&mall.Cart{})
	}
	return "删除成功！"
}

// Cartquery 查询购物车中相关商品的信息
func (m MallCart) Cartquery(uid string) []response.Goodsdata {
	temp := make([]response.Goodsdata, 0)
	global.GlobalDB.Raw("SELECT\n\tc.id cart_id, m.goods_id goods_id,\n\tm.specification_id specification_id,\n\tm.goods_name goods_name,\n\tm.coverimage cover_image,\n\tm.`specific` `specific`,\n\tm.price price,\n\tc.count count, \n\tm.color color \nFROM\n\tcarts c\n\tJOIN (\n\tSELECT\n\t\ta.goods_id goods_id,\n\t\ta.goods_name goods_name,\n\t\ta.goods_image_cover coverimage,\n\t\tb.`specific` `specific`,\n\t\tb.specification_id specification_id,\n\t\tb.color color,\n\t\tb.price price \n\tFROM\n\t\tgoods a\n\t\tJOIN specifications b ON a.goods_id = b.goods_id \n\t) m ON c.goods_id = m.goods_id \n\tAND c.specification_id = m.specification_id \n\tAND c.uid = ?", uid).Scan(&temp)
	res := make([]response.Goodsdata, 0)
	for _, v := range temp {
		v.CoverImage = "https://cdn.zmcicloud.cn/" + v.CoverImage
		res = append(res, v)
	}
	return res
}
