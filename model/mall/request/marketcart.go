package request

// AddCart 添加商品到购物车中
type AddCart struct {
	Count           int `json:"count"`           // 数量
	Goodsid         int `json:"goodsid"`         // 商品id
	Specificationid int `json:"specificationid"` // 商品规格id
}

// UpdateCount 更新购物车中商品的数量
type UpdateCount struct {
	Cartid          int `json:"cartid"`          // 购物车id
	Newcount        int `json:"newcount"`        // 新的数量
	Specificationid int `json:"specificationid"` // 商品规格id
}

// Cartdelete 删除购物车中的商品
type Cartdelete struct {
	Deletegoods []int `json:"deletegoods"` // Deletegoods
}

// CartQueryById 通过购物车id来查询商品
type CartQueryById struct {
	Cartids []int `json:"cartids"` // 需要查询的id组
}
