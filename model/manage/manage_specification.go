package manage

// 商品规格
type Specification struct {
	GoodsId         int     `json:"goodsId"`
	SpecificationId int     `json:"specificationId"`
	Color           string  `json:"color"`
	Specific        string  `json:"specific"`
	Count           int     `json:"count"`
	Price           float64 `json:"price"`
}
