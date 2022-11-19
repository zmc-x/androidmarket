package mall

// Orderitem 订单子项
type Orderitem struct {
	Id         int     `json:"id"`         // id
	Color      string  `json:"color"`      // 颜色
	Count      int     `json:"count"`      // 数量
	CoverImage string  `json:"coverImage"` // 商品封面
	GoodsName  string  `json:"goodsName"`  // 商品名称
	OrderId    int     `json:"OrderId"`    // 外键_订单id
	Price      float64 `json:"price"`      // 商品单价
	Specific   string  `json:"specific"`   // 商品规格
}
