package manage

import "androidmarket/service"

type ManageGroup struct {
	Managegoods // 管理商品信息
}

var AddGoodsMajorInfo = service.Service.ManageGroup.AddGoodsMajorInfo // 添加商品信息
var AddGoodsImage = service.Service.ManageGroup.AddGoodsImages        // 添加商品图片
var AddGoodsSpecific = service.Service.ManageGroup.AddGoodsSpecific   // 添加商品规格
var UpdateGoodsinfo = service.Service.ManageGroup.Updategoodsinfo     // 修改商品price || count
var DeleteGoodsinfo = service.Service.ManageGroup.Deletegoods         // 删除商品 or 商品的规格
