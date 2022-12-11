package mall

import (
	"androidmarket/global"
	"androidmarket/model/mall"
	"androidmarket/model/mall/response"
)

type MallGoods struct{}

// ShowGoodsInfo 返回商品的详情信息
func (m MallGoods) ShowGoodsInfo(goodsid int, specificationid int) (bool, response.ShowGoodsInfo) {
	temp := mall.Goods{}
	res := global.GlobalDB.Where("goods_id = ? and goods_status = 1", goodsid).Find(&temp)
	if res.RowsAffected == 0 {
		return false, response.ShowGoodsInfo{}
	} else {
		//	 处理图片path
		goodsinfo := response.ShowGoodsInfo{
			Goodsid:       goodsid,
			Goodsname:     temp.GoodsName,
			Coverimage:    "https://cdn.zmcicloud.cn/" + temp.GoodsImageCover,
			Specification: QuerySpecification(goodsid, specificationid),
		}
		imagelen := len(temp.GoodsImageInformation)
		cnt := 0
		for i, j := 0, 0; i < imagelen; i++ {
			for j = i + 1; j < imagelen; j++ {
				if temp.GoodsImageInformation[j] == ',' {
					str := "https://cdn.zmcicloud.cn/" + temp.GoodsImageInformation[i:j]
					goodsinfo.Images = append(goodsinfo.Images, str)
					break
				}
			}
			i = j
			cnt++
		}
		return true, goodsinfo
	}
}

// QueryGoodsByType 通过商品类别来查找相关的商品
func (m MallGoods) QueryGoodsByType(goostype string) (bool, []response.GoodsModel) {
	temp := make([]mall.Goods, 0)
	res := global.GlobalDB.Where("goods_type = ?", goostype).Find(&temp)
	if res.RowsAffected != 0 {
		result := make([]response.GoodsModel, 0)
		for _, v := range temp {
			result = append(result, response.GoodsModel{
				Goodsid:    v.GoodsId,
				Goodsname:  v.GoodsName,
				Goodscover: "https://cdn.zmcicloud.cn/" + v.GoodsImageCover,
			})
		}
		return true, result
	}
	return false, nil
}

// QuerySpecification 查询商品对应的规格
func QuerySpecification(goodsid int, specificationid int) response.Model {
	temp := mall.Specification{}
	res := global.GlobalDB.Where("goods_id = ? and specification_id = ?", goodsid, specificationid).Find(&temp)
	if res.RowsAffected != 0 {
		result := response.Model{
			Color:           temp.Color,
			Price:           temp.Price,
			Specific:        temp.Specific,
			Specificationid: temp.SpecificationId,
		}
		return result
	}
	return response.Model{}
}

// QueryHomeinfo 查询商品首页信息
func (m MallGoods) QueryHomeinfo() []response.GoodsHomeInfo {
	mid := make([]response.GoodsHomeInfo, 0)
	global.GlobalDB.Raw("SELECT\n\tm.goods_id goods_id,\n\tn.goods_name goods_name,\n\tm.specification_id specification_id,\n\tm.color color,\n\tm.price price,\n\tm.`specific` `specific`,\n\tn.goods_image_cover coverimage \nFROM\n\t(\n\tSELECT\n\t\ts.goods_id goods_id,\n\t\ts.`specific` `specific`,\n\t\ts.specification_id specification_id,\n\t\ts.color color,\n\t\ts.price price \n\tFROM\n\t\tspecifications s\n\t\tJOIN ( SELECT goods_id, min( price ) min_price FROM specifications GROUP BY goods_id ) b ON s.goods_id = b.goods_id \n\t\tAND s.price = b.min_price \n\t) m\n\tJOIN goods n \nWHERE\n\tm.goods_id = n.goods_id").Scan(&mid)
	res := make([]response.GoodsHomeInfo, 0)
	// 处理url
	for _, v := range mid {
		res = append(res, response.GoodsHomeInfo{
			GoodsId:         v.GoodsId,
			Specific:        v.Specific,
			SpecificationId: v.SpecificationId,
			Coverimage:      "https://cdn.zmcicloud.cn/" + v.Coverimage,
			GoodsName:       v.GoodsName,
			Price:           v.Price,
			Color:           v.Color,
		})
	}
	return res
}

// QueryGoodsinfo 生成订单时需要的相关商品信息
func (m MallGoods) QueryGoodsinfo(goodsid int, specificationid int) response.GoodsInOrderInfo {
	res := response.GoodsInOrderInfo{}
	global.GlobalDB.Model(&mall.Specification{}).Select("specifications.specification_id, goods.goods_id, goods.goods_name, specifications.specific, specifications.price, specifications.color, goods.goods_image_cover cover_image").Joins("join goods on goods.goods_id = specifications.goods_id").Where("specifications.specification_id = ? and specifications.goods_id = ?", specificationid, goodsid).Scan(&res)
	res.Count = 1
	res.CoverImage = "https://cdn.zmcicloud.cn/" + res.CoverImage
	return res
}

// QueryByName 通过商品名称来查询相关信息
func (m MallGoods) QueryByName(goodsname string) ([]response.GoodsHomeInfo, string, bool) {
	mid := make([]response.GoodsHomeInfo, 0)
	queryres := global.GlobalDB.Model(&mall.Goods{}).Select("specifications.specification_id, goods.goods_id, goods.goods_name, specifications.specific, specifications.price, specifications.color, goods.goods_image_cover coverimage").Joins("join specifications on goods.goods_id = specifications.goods_id").Where("goods.goods_name like ?", "%"+goodsname+"%").Find(&mid)
	res := make([]response.GoodsHomeInfo, 0)
	if queryres.RowsAffected == 0 {
		return res, "出错了哦，不存在这样的商品！", false
	}
	for _, v := range mid {
		res = append(res, response.GoodsHomeInfo{
			GoodsId:         v.GoodsId,
			Specific:        v.Specific,
			SpecificationId: v.SpecificationId,
			Coverimage:      "https://cdn.zmcicloud.cn/" + v.Coverimage,
			GoodsName:       v.GoodsName,
			Price:           v.Price,
			Color:           v.Color,
		})
	}
	return res, "查询成功！", true
}
