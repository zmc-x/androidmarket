package request

// AddCart 添加商品到购物车中
type AddCart struct {
	Count           int `json:"count"`           // 数量
	Goodsid         int `json:"goodsid"`         // 商品id
	Specificationid int `json:"specificationid"` // 商品规格id
}

// UpdateCount 更新购物车中商品的数量
type UpdateCount struct {
	Goodsid         int `json:"goodsid"`         // 商品id
	Newcount        int `json:"newcount"`        // 新的数量
	Specificationid int `json:"specificationid"` // 商品规格id
}

// Cartdelete 删除购物车中的商品
type Cartdelete struct {
	Deletegoods []Deletegood `json:"deletegoods"` // Deletegoods
}

type Deletegood struct {
	Goodsid         int `json:"goodsid"`         // 商品id
	Specificationid int `json:"specificationid"` // 规格id
}
