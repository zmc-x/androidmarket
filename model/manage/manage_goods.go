package manage

// Goods 商品信息
type Goods struct {
	GoodsId               int    `json:"goodsId"`
	GoodsName             string `json:"goodsName"`
	GoodsType             string `json:"goodsType"`
	GoodsImageCover       string `json:"goodsImageCover"`
	GoodsImageInformation string `json:"goodsImageInformation"`
	GoodsStatus           int    `json:"goodsStatus"`
}
