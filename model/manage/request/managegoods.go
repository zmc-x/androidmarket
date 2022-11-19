package request

// Goodsinfo 添加商品信息
type Goodsinfo struct {
	GoodsID   int    `json:"goodsId"`   // 商品id
	GoodsName string `json:"goodsName"` // 商品名称
	GoodsType string `json:"goodsType"` // 商品类别
}

// Specification 添加商品规格
type Specification struct {
	Count         int       `json:"count"`   // 数量
	Goodsid       int       `json:"goodsid"` // 商品id
	Price         float64   `json:"price"`   // 价格
	Specification Otherinfo `json:"specification"`
}

// Otherinfo 商品价格&数量
type Otherinfo struct {
	Color    string `json:"color"`    // 颜色
	Specific string `json:"specific"` // 规格
}

// Updateinfo 修改商品的信息
type Updateinfo struct {
	Count           int     `json:"count"`           // 数量
	Goodsid         int     `json:"goodsid"`         // 商品id
	Price           float64 `json:"price"`           // 价格
	Specificationid int     `json:"specificationid"` // 规格编号
}

// Deleteinfo 删除的商品 or 规格信息
type Deleteinfo struct {
	Deleteobject    Deleteobject `json:"deleteobject"`    // Deleteobject
	Goodsid         []int        `json:"goodsid"`         // 商品id组
	Specificationid []int        `json:"specificationid"` // 规格id组
}

// Deleteobject 选择那个对象来进行删除
type Deleteobject struct {
	Goods         bool `json:"goods"`         // 商品
	Specification bool `json:"specification"` // 商品规格
}
