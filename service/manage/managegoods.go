package manage

import (
	"androidmarket/global"
	"androidmarket/model/manage"
	"androidmarket/model/manage/request"
)

type ManageGoods struct{}

// AddGoodsMajorInfo 添加商品部分信息
func (m ManageGoods) AddGoodsMajorInfo(info request.Goodsinfo) (bool, error) {
	res := global.GlobalDB.Create(&manage.Goods{
		GoodsId:   info.GoodsID,
		GoodsName: info.GoodsName,
		GoodsType: info.GoodsType,
	})
	if res.Error != nil {
		return false, res.Error
	}
	return true, res.Error
}

// AddGoodsImages 添加商品图片信息
func (m ManageGoods) AddGoodsImages(cover string, info string, goodsid int) (bool, error) {
	res := global.GlobalDB.Model(&manage.Goods{}).Where("goods_id = ?", goodsid).Select("goods_image_cover", "goods_image_information").Updates(&manage.Goods{
		GoodsImageCover:       cover,
		GoodsImageInformation: info,
	})
	if res.Error != nil {
		return false, res.Error
	} else {
		return true, res.Error
	}
}

// AddGoodsSpecific 添加商品的规格&价格&数量
func (m ManageGoods) AddGoodsSpecific(param request.Specification) (string, bool, error) {
	res := global.GlobalDB.Create(&manage.Specification{
		GoodsId:  param.Goodsid,
		Color:    param.Specification.Color,
		Specific: param.Specification.Specific,
		Price:    param.Price,
		Count:    param.Count,
	})
	if res.Error != nil {
		return "添加失败", false, res.Error
	} else {
		if param.Count != 0 {
			// 修改上架信息
			global.GlobalDB.Model(&manage.Goods{}).Where("goods_id = ?", param.Goodsid).Update("goods_status", 1)
		}
		return "添加成功！", true, res.Error
	}
}

// Updategoodsinfo 修改商品的price && count
func (m ManageGoods) Updategoodsinfo(param request.Updateinfo) (string, bool, error) {
	// 开启事务
	tx := global.GlobalDB.Begin()
	res := tx.Model(&manage.Specification{}).Where("goods_id = ? and specification_id = ?", param.Goodsid, param.Specificationid).Select("count", "price").Updates(manage.Specification{
		Price: param.Price,
		Count: param.Count,
	})
	// 不存在该规格id
	if res.RowsAffected == 0 {
		// 事务回滚
		tx.Rollback()
		return "修改失败，非法修改", false, nil
	}
	if res.Error != nil {
		// 错误信息
		var msg = "修改失败!"
		// 事务回滚
		tx.Rollback()
		return msg, false, res.Error
	} else {
		// 查看是否该商品id下是否存在存货
		store := make([]manage.Specification, 0)
		tx.Where("goods_id = ?", param.Goodsid).Find(&store)
		// flag
		flag := true
		for _, v := range store {
			if v.Count != 0 {
				flag = false
				// 将上架信息改为1（上架
				tx.Model(&manage.Goods{}).Where("goods_id = ?", param.Goodsid).Update("goods_status", 1)
				break // 跳出循环
			}
		}
		if flag {
			// 将上架信息改为0（下架
			tx.Model(&manage.Goods{}).Where("goods_id = ?", param.Goodsid).Update("goods_status", 0)
		}
		// 提交事务
		tx.Commit()
		return "修改成功！", true, nil
	}
}

// Deletegoods 删除商品规格 or 商品
func (m ManageGoods) Deletegoods(param request.Deleteinfo) (string, bool) {
	// 开启事务
	tx := global.GlobalDB.Begin()
	if param.Deleteobject.Goods {
		// 删除商品
		for _, v := range param.Goodsid {
			r1 := tx.Where("goods_id = ?", v).Delete(&manage.Specification{})
			r2 := tx.Where("goods_id = ?", v).Delete(&manage.Goods{})
			if r1.Error != nil || r2.Error != nil {
				// 回滚
				tx.Rollback()
				return "删除商品出现错误！", false
			}
		}
		tx.Commit()
		return "删除商品成功！", true
	} else if param.Deleteobject.Specification {
		// 删除商品对应的规格
		for _, v := range param.Specificationid {
			res := tx.Where("goods_id = ? and specification_id = ?", param.Goodsid[0], v).Delete(&manage.Specification{})
			if res.Error != nil {
				tx.Rollback()
				return "删除商品规格错误！", false
			}
		}
		tx.Commit()
		return "删除商品规格成功！", true
	} else {
		tx.Rollback()
		return "删除对象未知！", false
	}
}
