package request

// Createoder 创建订单
type Createoder struct {
	Order     Order       `json:"order"`
	Orderitem []Orderitem `json:"orderitem"`
	Cartids   []int       `json:"cartids"`
}

type Order struct {
	Addressid int     `json:"addressid"` // 地址
	Allprice  float64 `json:"allprice"`  // 总价格
}

type Orderitem struct {
	Color      string  `json:"color"`      // 颜色
	Count      int     `json:"count"`      // 数量
	Coverimage string  `json:"coverimage"` // 商品封面
	Goodsname  string  `json:"goodsname"`  // 商品名称
	Price      float64 `json:"price"`      // 商品单价
	Specific   string  `json:"specific"`   // 商品规格
}
