package mall

import "androidmarket/service"

type MallGroup struct {
	Malluser    // 用户操作
	Malladdress // 地址操作
	MallGoods   // 商品操作
	MallCart    // 购物车操作
	Mallorder   // 订单操作
}

var Userlogin = service.Service.MallGroup.Userlogin               // 登录
var UserSignup = service.Service.MallGroup.Usersignup             // 注册
var UserUpdatePass = service.Service.MallGroup.Updatepass         // 修改用户密码
var UserInfo = service.Service.MallGroup.GetUserInfo              // 获取用户的相关信息
var Addaddress = service.Service.MallGroup.AddAddress             // 添加地址
var Updateaddress = service.Service.MallGroup.UpdateAddress       // 修改地址
var Deleteaddresses = service.Service.MallGroup.Deleteaddresses   // 删除地址
var Selectaddresses = service.Service.MallGroup.Selectaddresses   // 查询地址
var ShowGoodsInfo = service.Service.MallGroup.ShowGoodsInfo       // 查询单个商品的详细信息
var QueryGoodsByType = service.Service.MallGroup.QueryGoodsByType // 通过商品类型查询相关信息
var GoodsHomeInfo = service.Service.MallGroup.QueryHomeinfo       // 查询商城首页的信息
var QueryByName = service.Service.MallGroup.QueryByName           // 通过商品名称来进行查询
var Goodsinfo = service.Service.MallGroup.QueryGoodsinfo          // 查询商品在订单中的相关信息
var CartAdd = service.Service.MallGroup.Cartadd                   // 添加商品到购物车
var Cartupdate = service.Service.MallGroup.Cartupdatecount        // 修改购物车中商品的数量
var Cartdelete = service.Service.MallGroup.Cartdelete             // 删除购物车中的商品
var CartQuery = service.Service.MallGroup.Cartquery               // 查询购物车中商品信息
var CartQueryById = service.Service.MallGroup.CartQueryById       // 查询购物车中的商品信息by id
var OrderCreate = service.Service.MallGroup.CreateOrder           // 生成订单
var OrderUpdate = service.Service.MallGroup.Updateorder           // 修改订单的状态
var OrderQuery = service.Service.MallGroup.OrderQuery             // 查询订单
var OrderQueryInfo = service.Service.MallGroup.OrderQueryInfo     // 查询订单详细信息
