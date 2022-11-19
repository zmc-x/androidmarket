package mall

type MallRouteGroup struct {
	Malluser  // 用户route
	Address   // 地址route
	MallGoods // 商品route
	MallCart  // 购物车route
	MallOrder // 订单route
}
