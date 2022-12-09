package response

// ShowGoodsInfo 获取商品详情页信息
type ShowGoodsInfo struct {
	Coverimage    string   `json:"coverimage"`    // 商品封面path
	Goodsid       int      `json:"goodsid"`       // 商品id
	Goodsname     string   `json:"goodsname"`     // 商品名称
	Images        []string `json:"images"`        // 商品详情页
	Specification Model    `json:"specification"` // 商品规格
}

// Model 规格模型
type Model struct {
	Color           string  `json:"color"`           // 颜色
	Price           float64 `json:"price"`           // 价格
	Specific        string  `json:"specific"`        // 规格
	Specificationid int     `json:"specificationid"` // 规格id
}

// GoodsModel 商品详细信息模型
type GoodsModel struct {
	Goodscover string `json:"goodscover"` // 商品封面
	Goodsid    int    `json:"goodsid"`    // 商品id
	Goodsname  string `json:"goodsname"`  // 商品name
}

// GoodsHomeInfo 首页商品信息
type GoodsHomeInfo struct {
	GoodsId         int     `json:"goodsId"`         // 商品id
	SpecificationId int     `json:"specificationId"` // 商品规格id
	Specific        string  `json:"specific"`        // 商品规格
	Color           string  `json:"color"`           // 商品颜色
	Price           float64 `json:"price"`           // 商品价格
	GoodsName       string  `json:"goodsName"`       // 商品名称
	Coverimage      string  `json:"coverimage"`      // 商品封面
}

// GoodsInOrderInfo 商品信息在订单中的显示（商品直接购买的情况
type GoodsInOrderInfo struct {
	Color           string  `json:"color"`           // 颜色
	Count           int     `json:"count"`           // 商品数量
	CoverImage      string  `json:"coverImage"`      // 商品封面
	GoodsID         int     `json:"goodsId"`         // 商品id
	GoodsName       string  `json:"goodsName"`       // 商品名称
	Price           float64 `json:"price"`           // 价格
	Specific        string  `json:"specific"`        // 商品规格
	SpecificationID int     `json:"specificationId"` // 商品规格id
}
