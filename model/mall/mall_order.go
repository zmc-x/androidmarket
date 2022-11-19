package mall

// Order 订单
type Order struct {
	Id         int     `json:"id"`         // 订单id
	Uid        string  `json:"uid"`        // 用户id
	AddressId  int     `json:"addressId"`  // 地址
	Allprice   float64 `json:"allprice"`   // 总价格
	Createdat  string  `json:"createdat"`  // 创建时间
	Finishedat string  `json:"finishedat"` // 取消/完成时间
	Status     int     `json:"status"`     // 订单状态
}
