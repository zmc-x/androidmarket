package mall

type MallServiceGroup struct {
	Malluser    // 对user的相关操作
	MallAddress // 对地址表进行操作
	MallGoods   // 商品操作
	MallCart    // 购物车操作
	MallOrder   // 订单操作
}
