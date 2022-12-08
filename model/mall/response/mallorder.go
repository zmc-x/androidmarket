package response

// Orders 订单信息
type Orders struct {
	Goods       []Goodsinfo `json:"goods"`       // goodsinfo
	Orderid     int         `json:"orderid"`     // 订单id
	Allprice    float64     `json:"allprice"`    // 订单总价格
	Createdat   string      `json:"createdat"`   // 创建订单时间
	Orderstatus int         `json:"orderstatus"` // 订单的状态信息
}

// Orderinfo 订单详细信息
type Orderinfo struct {
	//Location   string      `json:"location"`   // 地址
	Province       string      `json:"province"`       // 省
	City           string      `json:"city"`           // 市
	County         string      `json:"county"`         // 区 / 县
	Detaillocation string      `json:"detaillocation"` // 详细地址
	Allprice       float64     `json:"allprice"`       // 总价格
	Createdat      string      `json:"createdat"`      // 创建时间
	Finishedat     string      `json:"finishedat"`     // 取消/完成时间
	Goodsinfo      []Goodsinfo `json:"goodsinfo"`      // GoodsInfo
	Status         int         `json:"status"`         // 订单状态
}

type Goodsinfo struct {
	Color      string  `json:"color"`      // 颜色
	Count      int     `json:"count"`      // 数量
	CoverImage string  `json:"coverImage"` // 商品封面
	GoodsName  string  `json:"goodsName"`  // 商品名称
	Price      float64 `json:"price"`      // 商品单价
	Specific   string  `json:"specific"`   // 商品规格
}
