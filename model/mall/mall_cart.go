package mall

type Cart struct {
	Id              int    `json:"id"`              // 购物车id
	Count           int    `json:"count"`           // 数量
	GoodsId         int    `json:"goodsId"`         // 商品id
	SpecificationId int    `json:"specificationId"` // 商品规格id
	Uid             string `json:"uid"`             // 用户id
}
