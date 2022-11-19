package request

type Address struct {
	Id             int    `json:"id"`
	Name           string `json:"name" binding:"required"`
	Location       string `json:"location" binding:"required"`
	Tel            int    `json:"tel" binding:"required"`
	Defaultaddress int    `json:"defaultaddress"`
}

// 删除需要的id组
type Addressesid struct {
	Ids []int `json:"ids"`
}
