package response

// ShowGoodsInfo 获取商品详情页信息
type ShowGoodsInfo struct {
	Coverimage    string   `json:"coverimage"`    // 商品封面path
	Goodsid       int      `json:"goodsid"`       // 商品id
	Goodsname     string   `json:"goodsname"`     // 商品名称
	Images        []string `json:"images"`        // 商品详情页
	Specification []Model  `json:"specification"` // 商品规格
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