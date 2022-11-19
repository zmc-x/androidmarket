package mall

import (
	"androidmarket/global"
	"androidmarket/model/mall"
	"androidmarket/model/mall/response"
)

type MallGoods struct{}

// ShowGoodsInfo 返回商品的详情信息
func (m MallGoods) ShowGoodsInfo(goodsid int) (bool, response.ShowGoodsInfo) {
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
			Specification: QuerySpecification(goodsid),
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
func QuerySpecification(goodsid int) []response.Model {
	temp := make([]mall.Specification, 0)
	res := global.GlobalDB.Where("goods_id = ?", goodsid).Find(&temp)
	if res.RowsAffected != 0 {
		result := make([]response.Model, 0)
		for _, v := range temp {
			result = append(result, response.Model{
				Specificationid: v.SpecificationId,
				Specific:        v.Specific,
				Price:           v.Price,
				Color:           v.Color,
			})
		}
		return result
	}
	return nil
}
