package mall

import (
	"androidmarket/global"
	"androidmarket/model/mall"
	"androidmarket/model/mall/request"
)

type MallAddress struct{}

// AddAddress 添加地址
func (m MallAddress) AddAddress(params request.Address, uid string) (string, bool, error) {
	// 判断同一个收件人下是否存在相同的信息
	temp := mall.Address{}
	midres := global.GlobalDB.Where("name = ? and location = ? and tel = ? and uid = ?", params.Name, params.Location, params.Tel, uid).Find(&temp)
	if midres.RowsAffected != 0 {
		return "添加失败，该收货地址已经存在", false, midres.Error
	}
	// 先判断该用户名下其是否需要添加默认地址
	if params.Defaultaddress == 1 {
		global.GlobalDB.Model(&mall.Address{}).Where("defaultaddress = ? and uid = ?", 1, uid).Update("defaultaddress", 0)
	}
	res := global.GlobalDB.Create(&mall.Address{Uid: uid, Name: params.Name, Tel: params.Tel, Defaultaddress: params.Defaultaddress, Location: params.Location})
	if res.Error == nil {
		// 添加成功
		return "Ok！收货地址添加成功", true, res.Error
	} else {
		// 出现错误
		return "Sorry！收货地址添加失败", false, res.Error
	}
}

// UpdateAddress 修改地址
func (m MallAddress) UpdateAddress(params request.Address, uid string) (string, bool, error) {
	formeraddress := mall.Address{}
	// 查找是否存在这条记录
	findres := global.GlobalDB.Where("id = ? and uid = ?", params.Id, uid).Find(&formeraddress)
	if findres.RowsAffected == 0 {
		return "该用户下不存在这条地址记录！", false, findres.Error
	}
	// 查看是否需要修改默认地址
	if params.Defaultaddress == 1 {
		global.GlobalDB.Where("uid = ?", formeraddress.Uid).Update("defaultaddress", 0)
	}
	res := global.GlobalDB.Model(&formeraddress).Select("name", "location", "tel", "defaultaddress").Updates(mall.Address{Name: params.Name, Location: params.Location, Tel: params.Tel, Defaultaddress: params.Defaultaddress})
	if res.Error != nil {
		return "Sorry，该地址记录修改失败", false, res.Error
	} else {
		return "ok， 该地址记录修改成功", true, res.Error
	}
}

// Deleteaddresses 删除地址
func (m MallAddress) Deleteaddresses(params []int, uid string) (string, bool) {
	res := global.GlobalDB.Where("uid = ?", uid).Delete(&mall.Address{}, params)
	if int(res.RowsAffected) < len(params) {
		return "Sorry，存在非法数据，无法删除！", false
	} else {
		return "Ok，地址删除成功！", true
	}
}

// Selectaddresses 查询地址
func (m MallAddress) Selectaddresses(uid string) []mall.Address {
	data := make([]mall.Address, 10)
	global.GlobalDB.Where("uid = ?", uid).Find(&data)
	return data
}
