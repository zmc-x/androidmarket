package request

// Login 登录信息&注册
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// updatepassword 修改用户密码
type Updatepass struct {
	Newpassword string `json:"newpassword" binding:"required"`
}
