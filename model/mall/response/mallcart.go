package response

type Goodsdata struct {
	Cartid          int     `json:"cartid"`          // 购物车的id
	Color           string  `json:"color"`           // 颜色
	Count           int     `json:"count"`           // 商品数量
	Coverimage      string  `json:"coverimage"`      // 商品封面
	GoodsID         int     `json:"goodsId"`         // 商品id
	GoodsName       string  `json:"goodsname"`       // 商品名称
	Price           float64 `json:"price"`           // 价格
	Specific        string  `json:"specific"`        // 商品规格
	SpecificationID int     `json:"specificationId"` // 商品规格id
}
